package syncer

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gogo/protobuf/types"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/memory"
	"github.com/solo-io/solo-kit/pkg/api/v1/reporter"
	"github.com/solo-io/solo-kit/pkg/utils/contextutils"
	"github.com/solo-io/solo-kit/pkg/utils/errutils"
	gloov1 "github.com/solo-io/solo-projects/projects/gloo/pkg/api/v1"
	"github.com/solo-io/solo-projects/projects/gloo/pkg/bootstrap"
	"github.com/solo-io/solo-projects/projects/gloo/pkg/defaults"
	"github.com/solo-io/solo-projects/projects/sqoop/pkg/api/v1"
	"github.com/solo-io/solo-projects/projects/sqoop/pkg/engine"
	"github.com/solo-io/solo-projects/projects/sqoop/pkg/engine/router"
	"github.com/solo-io/solo-projects/projects/sqoop/pkg/todo"
	"k8s.io/client-go/rest"
)

type Opts struct {
	WriteNamespace  string
	WatchNamespaces []string
	Schemas         factory.ResourceClientFactory
	ResolverMaps    factory.ResourceClientFactory
	Proxies         factory.ResourceClientFactory
	WatchOpts       clients.WatchOpts
	DevMode         bool
	SidecarAddr     string
}

func Setup(ctx context.Context, kubeCache *kube.KubeCache, cache memory.InMemoryResourceCache, settings *gloov1.Settings) error {
	var (
		cfg *rest.Config
	)
	proxyFactory, err := bootstrap.ConfigFactoryForSettings(
		settings,
		cache,
		kubeCache,
		gloov1.ProxyCrd,
		&cfg,
	)
	if err != nil {
		return err
	}

	schemaFactory, err := bootstrap.ConfigFactoryForSettings(
		settings,
		cache,
		kubeCache,
		v1.SchemaCrd,
		&cfg,
	)
	if err != nil {
		return err
	}

	resolverMapFactory, err := bootstrap.ConfigFactoryForSettings(
		settings,
		cache,
		kubeCache,
		v1.ResolverMapCrd,
		&cfg,
	)
	if err != nil {
		return err
	}

	refreshRate, err := types.DurationFromProto(settings.RefreshRate)
	if err != nil {
		return err
	}

	writeNamespace := settings.DiscoveryNamespace
	if writeNamespace == "" {
		writeNamespace = defaults.GlooSystem
	}
	watchNamespaces := settings.WatchNamespaces
	var writeNamespaceProvided bool
	for _, ns := range watchNamespaces {
		if ns == writeNamespace {
			writeNamespaceProvided = true
			break
		}
	}
	if !writeNamespaceProvided {
		watchNamespaces = append(watchNamespaces, writeNamespace)
	}
	opts := Opts{
		WriteNamespace:  writeNamespace,
		WatchNamespaces: watchNamespaces,
		Schemas:         schemaFactory,
		ResolverMaps:    resolverMapFactory,
		Proxies:         proxyFactory,
		WatchOpts: clients.WatchOpts{
			Ctx:         ctx,
			RefreshRate: refreshRate,
		},
		SidecarAddr: fmt.Sprintf("%v:%v", "127.0.0.1", TODO.SqoopSidecarBindPort),
	}
	return RunSqoop(opts)
}

func RunSqoop(opts Opts) error {
	watchNamespaces := opts.WatchNamespaces
	opts.WatchOpts = opts.WatchOpts.WithDefaults()
	opts.WatchOpts.Ctx = contextutils.WithLogger(opts.WatchOpts.Ctx, "gateway")

	// TODO(ilackarms): this piece (initalizing clients) should really be generated by solo-kit
	proxyClient, err := gloov1.NewProxyClient(opts.Proxies)
	if err != nil {
		return err
	}
	if err := proxyClient.Register(); err != nil {
		return err
	}
	proxyReconciler := gloov1.NewProxyReconciler(proxyClient)

	schemaClient, err := v1.NewSchemaClient(opts.Schemas)
	if err != nil {
		return err
	}
	if err := schemaClient.Register(); err != nil {
		return err
	}

	resolverMapClient, err := v1.NewResolverMapClient(opts.ResolverMaps)
	if err != nil {
		return err
	}
	if err := resolverMapClient.Register(); err != nil {
		return err
	}

	emitter := v1.NewApiEmitter(resolverMapClient, schemaClient)

	rpt := reporter.NewReporter("sqoop", resolverMapClient.BaseClient(), schemaClient.BaseClient())
	writeErrs := make(chan error)
	/*
		proxyReconciler:   proxyReconciler,
		engine:            engine,
		router:            router,
	*/
	eng := engine.NewEngine(opts.SidecarAddr)

	rtr := router.NewRouter()

	sync := NewGraphQLSyncer(opts.WriteNamespace, rpt, writeErrs, proxyReconciler, resolverMapClient, eng, rtr)

	go func() {
		contextutils.LoggerFrom(opts.WatchOpts.Ctx).Fatalf("failed starting sqoop server: %v",
			http.ListenAndServe(fmt.Sprintf(":%v", TODO.SqoopServerBindPort), rtr))
	}()

	eventLoop := v1.NewApiEventLoop(emitter, sync)
	eventLoopErrs, err := eventLoop.Run(watchNamespaces, opts.WatchOpts)
	if err != nil {
		return err
	}
	go errutils.AggregateErrs(opts.WatchOpts.Ctx, writeErrs, eventLoopErrs, "event_loop")

	logger := contextutils.LoggerFrom(opts.WatchOpts.Ctx)

	go func() {
		for {
			select {
			case err := <-writeErrs:
				logger.Errorf("error: %v", err)
			case <-opts.WatchOpts.Ctx.Done():
				close(writeErrs)
				return
			}
		}
	}()
	return nil
}
