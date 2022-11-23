package extauth

import (
	"context"
	"fmt"
	"sort"

	"github.com/solo-io/solo-projects/projects/extauth/pkg/runner"
	"go.uber.org/zap"

	errors "github.com/rotisserie/eris"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"

	"github.com/golang/protobuf/proto"
	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	extauth "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/extauth/v1"
	gloov1snap "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/gloosnapshot"
	"github.com/solo-io/gloo/projects/gloo/pkg/syncer"
	glooutils "github.com/solo-io/gloo/projects/gloo/pkg/utils"
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/go-utils/hashutils"
	envoycache "github.com/solo-io/solo-kit/pkg/api/v1/control-plane/cache"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v2/reporter"
	extAuthPlugin "github.com/solo-io/solo-projects/projects/gloo/pkg/plugins/extauth"
	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
)

// Compile-time assertion
var (
	_ syncer.TranslatorSyncerExtension = new(translatorSyncerExtension)
)

var (
	extauthConnectedStateDescription = "zero indicates gloo detected an error with the auth config and did not update its XDS snapshot, check the gloo logs for errors"
	extauthConnectedState            = stats.Int64("glooe.extauth/connected_state", extauthConnectedStateDescription, "1")

	extauthConnectedStateView = &view.View{
		Name:        "glooe.extauth/connected_state",
		Measure:     extauthConnectedState,
		Description: extauthConnectedStateDescription,
		Aggregation: view.LastValue(),
		TagKeys:     []tag.Key{},
	}
)

var (
	emptyTypedResources = map[string]envoycache.Resources{
		extauth.ExtAuthConfigType: {
			Version: "empty",
			Items:   map[string]envoycache.Resource{},
		},
	}
)

func init() {
	_ = view.Register(extauthConnectedStateView)
}

type translatorSyncerExtension struct {
	hasher func(resources []envoycache.Resource) (uint64, error)
}

func NewTranslatorSyncerExtension(_ context.Context, params syncer.TranslatorSyncerExtensionParams) syncer.TranslatorSyncerExtension {
	return &translatorSyncerExtension{
		hasher: params.Hasher,
	}
}

func (s *translatorSyncerExtension) ID() string {
	return runner.ServerRole
}

func (s *translatorSyncerExtension) Sync(
	ctx context.Context,
	snap *gloov1snap.ApiSnapshot,
	settings *gloov1.Settings,
	snapshotSetter syncer.SnapshotSetter,
	reports reporter.ResourceReports,
) {
	ctx = contextutils.WithLogger(ctx, "extAuthTranslatorSyncer")
	logger := contextutils.LoggerFrom(ctx)
	// only hash during debug because of the performance issues surrounding hashing
	if contextutils.GetLogLevel() == zap.DebugLevel {
		snapHash := hashutils.MustHash(snap)
		logger.Debugf("begin auth sync %v (%v proxies, %v upstreams, %v endpoints, %v secrets, %v artifacts, %v auth configs)", snapHash,
			len(snap.Proxies), len(snap.Upstreams), len(snap.Endpoints), len(snap.Secrets), len(snap.Artifacts), len(snap.AuthConfigs))
		defer logger.Debugf("end auth sync %v", snapHash)
	}

	s.SyncAndSet(ctx, snap, settings, snapshotSetter, reports)
}

func (s *translatorSyncerExtension) SyncAndSet(
	ctx context.Context,
	snap *gloov1snap.ApiSnapshot,
	settings *gloov1.Settings,
	snapshotSetter syncer.SnapshotSetter,
	reports reporter.ResourceReports,
) {
	helper := newHelper()
	reports.Accept(snap.AuthConfigs.AsInputResources()...)
	reports.Accept(snap.Proxies.AsInputResources()...)

	for _, proxy := range snap.Proxies {
		for _, listener := range proxy.Listeners {
			virtualHosts := glooutils.GetVirtualHostsForListener(listener)

			for _, virtualHost := range virtualHosts {
				virtualHost = proto.Clone(virtualHost).(*gloov1.VirtualHost)
				virtualHost.Name = glooutils.SanitizeForEnvoy(ctx, virtualHost.Name, "virtual host")

				if err := helper.processAuthExtension(ctx, snap, settings, virtualHost.GetOptions().GetExtauth(), reports, proxy); err != nil {
					// Continue to next virtualHost, error has been added to the report.
					continue
				}

				for _, route := range virtualHost.Routes {
					if err := helper.processAuthExtension(ctx, snap, settings, route.GetOptions().GetExtauth(), reports, proxy); err != nil {
						// Continue to next route, error has been added to the report.
						continue
					}

					for _, weightedDestination := range route.GetRouteAction().GetMulti().GetDestinations() {
						if err := helper.processAuthExtension(ctx, snap, settings, weightedDestination.GetOptions().GetExtauth(),
							reports, proxy); err != nil {
							// Continue to next weighted destination, error has been added to the report.
							continue
						}
					}
				}
			}
		}
	}

	var snapshotResources []envoycache.Resource
	for _, cfg := range ConvertConfigMapToSortedList(helper.translatedConfigs) {
		resource := extauth.NewExtAuthConfigXdsResourceWrapper(cfg)
		snapshotResources = append(snapshotResources, resource)
	}

	var extAuthSnapshot envoycache.Snapshot
	snapshotHash, err := s.hasher(snapshotResources)
	if err != nil {
		contextutils.LoggerFrom(ctx).DPanicw("error trying to hash snapshot resources for extauth translation", err)
	}

	if snapshotResources == nil || err != nil {
		// If there are no correctly formatted auth configs, use an empty configuration
		//
		// The SnapshotCache can now differentiate between nil and empty snapshotResources in a snapshot.
		// This was introduced with: https://github.com/solo-io/solo-kit/pull/410
		// A nil resource is not updated, whereas an empty resource is intended to be modified.
		//
		// The extauth service only becomes healthy after it has received auth configuration
		// from Gloo via xDS. Therefore, we must set the auth config resource to empty in the snapshot
		// so that extauth picks up the empty config, and becomes healthy
		extAuthSnapshot = envoycache.NewGenericSnapshot(emptyTypedResources)
	} else {
		snapshotVersion := fmt.Sprintf("%d", snapshotHash)
		extAuthSnapshot = envoycache.NewEasyGenericSnapshot(snapshotVersion, snapshotResources)
	}

	snapshotSetter.SetSnapshot(s.ID(), extAuthSnapshot)
	stats.Record(ctx, extauthConnectedState.M(int64(1)))
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

func (h *helper) processAuthExtension(ctx context.Context, snap *gloov1snap.ApiSnapshot, settings *gloov1.Settings, config *extauth.ExtAuthExtension,
	reports reporter.ResourceReports, parentProxy resources.InputResource) error {
	if config.GetConfigRef() != nil {
		return h.processAuthExtensionConfigRef(ctx, snap, config.GetConfigRef(), reports, parentProxy)
	}

	if config.GetCustomAuth() != nil {
		return h.processAuthExtensionCustomAuth(ctx, settings, config.GetCustomAuth(), reports, parentProxy)
	}

	// Just return if there is nothing to process
	return nil
}

func (h *helper) processAuthExtensionConfigRef(ctx context.Context, snap *gloov1snap.ApiSnapshot, configRef *core.ResourceRef,
	reports reporter.ResourceReports, parentProxy resources.InputResource) error {

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
		reports.AddError(parentProxy, err)
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

func (h *helper) processAuthExtensionCustomAuth(ctx context.Context, settings *gloov1.Settings, customAuth *extauth.CustomAuth,
	reports reporter.ResourceReports, parentProxy resources.InputResource) error {

	customAuthServerName := customAuth.GetName()
	if customAuthServerName == "" {
		// If name is not specified, there is nothing to validate
		return nil
	}

	namedExtAuthSettings := settings.GetNamedExtauth()
	if namedExtAuthSettings == nil {
		// A name is specified, but no settings are configured
		err := errors.New("Unable to find named_extauth in Settings")
		contextutils.LoggerFrom(ctx).Warnf("%v", err)
		reports.AddError(parentProxy, err)
		return err
	}

	if _, ok := namedExtAuthSettings[customAuthServerName]; !ok {
		// A name is specified, but it isn't one of the settings that are configured
		err := errors.Errorf("Unable to find custom auth server [%s] in named_extauth in Settings", customAuthServerName)
		contextutils.LoggerFrom(ctx).Warnf("%v", err)
		reports.AddError(parentProxy, err)
		return err
	}

	return nil
}

// We need a stable ordering of configs. It doesn't really matter what order as long as it's consistent.
// If we don't do this, different orders of configs will produce different
// hashes, which will trick the system into unnecessarily thinking that we need to update the extauth service.
// Visible for testing
func ConvertConfigMapToSortedList(configMap map[string]*extauth.ExtAuthConfig) []*extauth.ExtAuthConfig {
	// extract values for sorting
	var configs []*extauth.ExtAuthConfig
	var configKeys []string
	for key, cfg := range configMap {
		configs = append(configs, cfg)
		configKeys = append(configKeys, key)
	}

	sort.SliceStable(configs, func(i, j int) bool {
		return configKeys[i] < configKeys[j]
	})

	return configs
}
