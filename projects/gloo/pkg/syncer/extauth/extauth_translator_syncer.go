package extauth

import (
	"context"
	"fmt"

	"github.com/gogo/protobuf/proto"
	"github.com/mitchellh/hashstructure"
	errors "github.com/rotisserie/eris"
	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	extauth "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/extauth/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/syncer"
	glooutils "github.com/solo-io/gloo/projects/gloo/pkg/utils"
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/go-utils/hashutils"
	envoycache "github.com/solo-io/solo-kit/pkg/api/v1/control-plane/cache"
	"github.com/solo-io/solo-kit/pkg/api/v2/reporter"
	"github.com/solo-io/solo-projects/projects/extauth/pkg/runner"
	extAuthPlugin "github.com/solo-io/solo-projects/projects/gloo/pkg/plugins/extauth"
	"go.uber.org/zap"
)

type ExtAuthTranslatorSyncerExtension struct {
	reporter reporter.Reporter
}

func NewTranslatorSyncerExtension(params syncer.TranslatorSyncerExtensionParams) *ExtAuthTranslatorSyncerExtension {
	return &ExtAuthTranslatorSyncerExtension{reporter: params.Reporter}
}

// TODO(marco): report errors on auth config resources once we have the strongly typed API. Currently it is not possible
//  to do this consistently, since we need to parse the raw extension to get to the auth config, an operation that might itself fail.
func (s *ExtAuthTranslatorSyncerExtension) Sync(ctx context.Context, snap *gloov1.ApiSnapshot, xdsCache envoycache.SnapshotCache) error {
	ctx = contextutils.WithLogger(ctx, "extAuthTranslatorSyncer")
	logger := contextutils.LoggerFrom(ctx)
	snapHash := hashutils.MustHash(snap)
	logger.Infof("begin auth sync %v (%v proxies, %v upstreams, %v endpoints, %v secrets, %v artifacts, %v auth configs)", snapHash,
		len(snap.Proxies), len(snap.Upstreams), len(snap.Endpoints), len(snap.Secrets), len(snap.Artifacts), len(snap.AuthConfigs))
	defer logger.Infof("end auth sync %v", snapHash)

	return s.SyncAndSet(ctx, snap, xdsCache)
}

type SnapshotSetter interface {
	SetSnapshot(node string, snapshot envoycache.Snapshot) error
}

func (s *ExtAuthTranslatorSyncerExtension) SyncAndSet(ctx context.Context, snap *gloov1.ApiSnapshot, xdsCache SnapshotSetter) error {
	helper := newHelper()
	reports := make(reporter.ResourceReports)
	reports.Accept(snap.AuthConfigs.AsInputResources()...)

	for _, proxy := range snap.Proxies {
		for _, listener := range proxy.Listeners {
			httpListener, ok := listener.ListenerType.(*gloov1.Listener_HttpListener)
			if !ok {
				// not an http listener - skip it as currently ext auth is only supported for http
				continue
			}

			virtualHosts := httpListener.HttpListener.VirtualHosts
			for _, virtualHost := range virtualHosts {

				virtualHost = proto.Clone(virtualHost).(*gloov1.VirtualHost)
				virtualHost.Name = glooutils.SanitizeForEnvoy(ctx, virtualHost.Name, "virtual host")

				if err := helper.processAuthExtension(ctx, snap, virtualHost.GetOptions().GetExtauth(), reports); err != nil {
					return err
				}

				for _, route := range virtualHost.Routes {

					if err := helper.processAuthExtension(ctx, snap, route.GetOptions().GetExtauth(), reports); err != nil {
						return err
					}

					for _, weightedDestination := range route.GetRouteAction().GetMulti().GetDestinations() {
						if err := helper.processAuthExtension(ctx, snap, weightedDestination.GetOptions().GetExtauth(),
							reports); err != nil {
							return err
						}
					}
				}
			}
		}
	}

	var resources []envoycache.Resource
	for _, cfg := range helper.translatedConfigs {
		resource := extauth.NewExtAuthConfigXdsResourceWrapper(cfg)
		resources = append(resources, resource)
	}
	h, err := hashstructure.Hash(resources, nil)
	if err != nil {
		contextutils.LoggerFrom(ctx).With(zap.Error(err)).DPanic("error hashing ext auth")
		return err
	}
	extAuthSnapshot := envoycache.NewEasyGenericSnapshot(fmt.Sprintf("%d", h), resources)
	_ = xdsCache.SetSnapshot(runner.ExtAuthServerRole, extAuthSnapshot)
	if err := s.reporter.WriteReports(ctx, reports, nil); err != nil {
		contextutils.LoggerFrom(ctx).Warnf("Failed writing report for auth configs: %v", err)
		return errors.Wrapf(err, "writing reports")
	}
	return nil
}

// This translation helper contains a map where each key is the unique identifier of an AuthConfig and the corresponding
// value is the translated config. We use it avoid translating the same configuration multiple times.
type helper struct {
	translatedConfigs map[string]*extauth.ExtAuthConfig
}

func newHelper() *helper {
	return &helper{
		translatedConfigs: make(map[string]*extauth.ExtAuthConfig),
	}
}

func (h *helper) processAuthExtension(ctx context.Context, snap *gloov1.ApiSnapshot, config *extauth.ExtAuthExtension,
	reports reporter.ResourceReports) error {
	configRef := config.GetConfigRef()
	if configRef == nil {
		// Just return if there is nothing to translate
		return nil
	}

	// Don't perform duplicate work if we already have translated this resource
	if _, ok := h.translatedConfigs[configRef.Key()]; ok {
		return nil
	}

	cfg, err := snap.AuthConfigs.Find(configRef.GetNamespace(), configRef.GetName())
	if err != nil {
		contextutils.LoggerFrom(ctx).Warnf("Unable to find referenced auth config %v in snapshot", configRef)
		return err
	}

	// do validation
	extAuthPlugin.ValidateAuthConfig(cfg, reports)

	translatedConfig, err := extAuthPlugin.TranslateExtAuthConfig(ctx, snap, configRef)
	if err != nil {
		reports.AddError(cfg, err)
		return err
	} else if translatedConfig == nil {
		// Do nothing if config is empty or consists only of custom auth configs
		return nil
	}

	h.translatedConfigs[configRef.Key()] = translatedConfig
	return nil
}
