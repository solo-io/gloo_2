package proxy_syncer

import (
	"context"

	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	apiv1 "sigs.k8s.io/gateway-api/apis/v1"

	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/utils/statusutils"

	"github.com/solo-io/gloo/projects/gateway2/extensions"
	"github.com/solo-io/gloo/projects/gateway2/query"
	"github.com/solo-io/gloo/projects/gateway2/reports"
	gwv2_translator "github.com/solo-io/gloo/projects/gateway2/translator"
	gwplugins "github.com/solo-io/gloo/projects/gateway2/translator/plugins"
	"github.com/solo-io/gloo/projects/gateway2/translator/plugins/registry"
	gloo_solo_io "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/translator"
	"github.com/solo-io/gloo/projects/gloo/pkg/utils"
)

type ProxySyncer struct {
	translator     translator.Translator
	controllerName string
	writeNamespace string

	inputs          *GatewayInputChannels
	mgr             manager.Manager
	k8sGwExtensions extensions.K8sGatewayExtensions

	// proxyReconciler wraps the client that writes Proxy resources into an in-memory cache
	// This cache is utilized by the debug.ProxyEndpointServer
	proxyReconciler gloo_solo_io.ProxyReconciler
}

type GatewayInputChannels struct {
	genericEvent AsyncQueue[struct{}]
	secretEvent  AsyncQueue[SecretInputs]
}

func (x *GatewayInputChannels) Kick(ctx context.Context) {
	x.genericEvent.Enqueue(struct{}{})
}

func (x *GatewayInputChannels) UpdateSecretInputs(ctx context.Context, inputs SecretInputs) {
	x.secretEvent.Enqueue(inputs)
}

func NewGatewayInputChannels() *GatewayInputChannels {
	return &GatewayInputChannels{
		genericEvent: NewAsyncQueue[struct{}](),
		secretEvent:  NewAsyncQueue[SecretInputs](),
	}
}

func NewProxySyncer(
	controllerName, writeNamespace string,
	translator translator.Translator,
	inputs *GatewayInputChannels,
	mgr manager.Manager,
	k8sGwExtensions extensions.K8sGatewayExtensions,
	proxyClient gloo_solo_io.ProxyClient,
) *ProxySyncer {
	return &ProxySyncer{
		controllerName:  controllerName,
		writeNamespace:  writeNamespace,
		translator:      translator,
		inputs:          inputs,
		mgr:             mgr,
		k8sGwExtensions: k8sGwExtensions,
		proxyReconciler: gloo_solo_io.NewProxyReconciler(proxyClient, statusutils.NewNoOpStatusClient()),
	}
}

func (s *ProxySyncer) Start(ctx context.Context) error {
	ctx = contextutils.WithLogger(ctx, "k8s-gw-syncer")
	contextutils.LoggerFrom(ctx).Debug("starting syncer for k8s gateway proxies")

	var (
		secretsWarmed bool
	)
	resyncProxies := func() {
		if !secretsWarmed {
			return
		}
		contextutils.LoggerFrom(ctx).Debug("resyncing k8s gateway proxies")

		var gwl apiv1.GatewayList
		err := s.mgr.GetClient().List(ctx, &gwl)
		if err != nil {
			// This should never happen, try again?
			return
		}

		gatewayQueries := query.NewData(s.mgr.GetClient(), s.mgr.GetScheme())

		pluginRegistry := s.k8sGwExtensions.CreatePluginRegistry(ctx)
		gatewayTranslator := gwv2_translator.NewTranslator(gatewayQueries, pluginRegistry)

		rm := reports.NewReportMap()
		r := reports.NewReporter(&rm)

		var (
			proxies            gloo_solo_io.ProxyList
			translatedGateways []gwplugins.TranslatedGateway
		)
		for _, gw := range gwl.Items {
			proxy := gatewayTranslator.TranslateProxy(ctx, &gw, s.writeNamespace, r)
			if proxy != nil {
				proxies = append(proxies, proxy)
				translatedGateways = append(translatedGateways, gwplugins.TranslatedGateway{
					Gateway: gw,
				})
				//TODO: handle reports and process statuses
			}
		}

		applyPostTranslationPlugins(ctx, pluginRegistry, &gwplugins.PostTranslationContext{
			TranslatedGateways: translatedGateways,
		})

		s.syncStatus(ctx, rm, gwl)
		s.syncRouteStatus(ctx, rm)
		s.reconcileProxies(ctx, proxies)
	}

	for {
		select {
		case <-ctx.Done():
			contextutils.LoggerFrom(ctx).Debug("context done, stopping proxy syncer")
			return nil
		case <-s.inputs.genericEvent.Next():
			resyncProxies()
		case <-s.inputs.secretEvent.Next():
			secretsWarmed = true
			resyncProxies()
		}
	}
}

func (s *ProxySyncer) syncRouteStatus(ctx context.Context, rm reports.ReportMap) {
	ctx = contextutils.WithLogger(ctx, "routeStatusSyncer")
	logger := contextutils.LoggerFrom(ctx)
	rl := apiv1.HTTPRouteList{}
	err := s.mgr.GetClient().List(ctx, &rl)
	if err != nil {
		logger.Error(err)
		return
	}

	for _, route := range rl.Items {
		route := route // pike
		if status := rm.BuildRouteStatus(ctx, route, s.controllerName); status != nil {
			route.Status = *status
			if err := s.mgr.GetClient().Status().Update(ctx, &route); err != nil {
				logger.Error(err)
			}
		}
	}
}

// syncStatus updates the status of the Gateway CRs
func (s *ProxySyncer) syncStatus(ctx context.Context, rm reports.ReportMap, gwl apiv1.GatewayList) {
	ctx = contextutils.WithLogger(ctx, "statusSyncer")
	logger := contextutils.LoggerFrom(ctx)
	for _, gw := range gwl.Items {
		gw := gw // pike
		if status := rm.BuildGWStatus(ctx, gw); status != nil {
			gw.Status = *status
			if err := s.mgr.GetClient().Status().Patch(ctx, &gw, client.Merge); err != nil {
				logger.Error(err)
			}
		}
	}
}

// reconcileProxies persists the proxies that were generated during translations and stores them in an in-memory cache
// This cache is utilized by the debug.ProxyEndpointServer
// As well as to resync the Gloo Xds Translator (when it receives new proxies using a MultiResourceClient)
func (s *ProxySyncer) reconcileProxies(ctx context.Context, proxyList gloo_solo_io.ProxyList) {
	ctx = contextutils.WithLogger(ctx, "proxyCache")
	logger := contextutils.LoggerFrom(ctx)

	// Proxy CR is located in the same namespace as the originating Gateway CR
	// As a result, we may have a list of Proxies that are in different namespaces
	// Since the reconciler accepts the namespace as an argument, we need to split
	// the list so we have a lists of proxies, isolated to each namespace
	//var proxyListByNamespace = make(map[string]gloo_solo_io.ProxyList)
	//
	//for _, proxy := range proxyList {
	//	proxyNs := proxy.GetMetadata().GetNamespace()
	//	nsList, ok := proxyListByNamespace[proxyNs]
	//	if ok {
	//		nsList = append(nsList, proxy)
	//		proxyListByNamespace[proxyNs] = nsList
	//	} else {
	//		proxyListByNamespace[proxyNs] = gloo_solo_io.ProxyList{
	//			proxy,
	//		}
	//	}
	//}

	//for ns, nsList := range proxyListByNamespace {
	err := s.proxyReconciler.Reconcile(
		s.writeNamespace, //ns,
		proxyList,        //nsList,
		func(original, desired *gloo_solo_io.Proxy) (bool, error) {
			// ignore proxies that do not have our owner label
			if original.GetMetadata().GetLabels() == nil || original.GetMetadata().GetLabels()[utils.ProxyTypeKey] != utils.GlooGatewayProxyValue {
				// TODO(npolshak): Currently we update all Gloo Gateway proxies. We should create a new label and ignore proxies that are not owned by Gloo control plane running in a specific namespace via POD_NAMESPACE
				logger.Debugf("ignoring proxy %v in namespace %v, does not have owner label %v", original.GetMetadata().GetName(), original.GetMetadata().GetNamespace(), utils.GlooGatewayProxyValue)
				return false, nil
			}
			// otherwise always update
			return true, nil
		},
		clients.ListOpts{
			Ctx: ctx,
		})
	if err != nil {
		// A write error to our cache should not impact translation
		// We will emit a message, and continue
		logger.Error(err)
	}
	//}
}

func applyPostTranslationPlugins(ctx context.Context, pluginRegistry registry.PluginRegistry, translationContext *gwplugins.PostTranslationContext) {
	ctx = contextutils.WithLogger(ctx, "postTranslation")
	logger := contextutils.LoggerFrom(ctx)

	for _, postTranslationPlugin := range pluginRegistry.GetPostTranslationPlugins() {
		err := postTranslationPlugin.ApplyPostTranslationPlugin(ctx, translationContext)
		if err != nil {
			logger.Errorf("Error applying post-translation plugin: %v", err)
			continue
		}
	}
}
