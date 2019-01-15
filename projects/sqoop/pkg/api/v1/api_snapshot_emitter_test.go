// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"context"
	"os"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/go-utils/kubeutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"
	kuberc "github.com/solo-io/solo-kit/pkg/api/v1/clients/kube"
	"github.com/solo-io/solo-kit/pkg/utils/log"
	"github.com/solo-io/solo-kit/test/helpers"
	"github.com/solo-io/solo-kit/test/setup"
	"k8s.io/client-go/rest"

	// Needed to run tests in GKE
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"

	// From https://github.com/kubernetes/client-go/blob/53c7adfd0294caa142d961e1f780f74081d5b15f/examples/out-of-cluster-client-configuration/main.go#L31
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

var _ = Describe("V1Emitter", func() {
	if os.Getenv("RUN_KUBE_TESTS") != "1" {
		log.Printf("This test creates kubernetes resources and is disabled by default. To enable, set RUN_KUBE_TESTS=1 in your env.")
		return
	}
	var (
		namespace1        string
		namespace2        string
		name1, name2      = "angela" + helpers.RandString(3), "bob" + helpers.RandString(3)
		cfg               *rest.Config
		emitter           ApiEmitter
		resolverMapClient ResolverMapClient
		schemaClient      SchemaClient
	)

	BeforeEach(func() {
		namespace1 = helpers.RandString(8)
		namespace2 = helpers.RandString(8)
		var err error
		cfg, err = kubeutils.GetConfig("", "")
		Expect(err).NotTo(HaveOccurred())
		err = setup.SetupKubeForTest(namespace1)
		Expect(err).NotTo(HaveOccurred())
		err = setup.SetupKubeForTest(namespace2)
		Expect(err).NotTo(HaveOccurred())
		// ResolverMap Constructor
		resolverMapClientFactory := &factory.KubeResourceClientFactory{
			Crd:         ResolverMapCrd,
			Cfg:         cfg,
			SharedCache: kuberc.NewKubeCache(),
		}
		resolverMapClient, err = NewResolverMapClient(resolverMapClientFactory)
		Expect(err).NotTo(HaveOccurred())
		// Schema Constructor
		schemaClientFactory := &factory.KubeResourceClientFactory{
			Crd:         SchemaCrd,
			Cfg:         cfg,
			SharedCache: kuberc.NewKubeCache(),
		}
		schemaClient, err = NewSchemaClient(schemaClientFactory)
		Expect(err).NotTo(HaveOccurred())
		emitter = NewApiEmitter(resolverMapClient, schemaClient)
	})
	AfterEach(func() {
		setup.TeardownKube(namespace1)
		setup.TeardownKube(namespace2)
	})
	It("tracks snapshots on changes to any resource", func() {
		ctx := context.Background()
		err := emitter.Register()
		Expect(err).NotTo(HaveOccurred())

		snapshots, errs, err := emitter.Snapshots([]string{namespace1, namespace2}, clients.WatchOpts{
			Ctx:         ctx,
			RefreshRate: time.Second,
		})
		Expect(err).NotTo(HaveOccurred())

		var snap *ApiSnapshot

		/*
			ResolverMap
		*/

		assertSnapshotResolverMaps := func(expectResolverMaps ResolverMapList, unexpectResolverMaps ResolverMapList) {
		drain:
			for {
				select {
				case snap = <-snapshots:
					for _, expected := range expectResolverMaps {
						if _, err := snap.ResolverMaps.List().Find(expected.Metadata.Ref().Strings()); err != nil {
							continue drain
						}
					}
					for _, unexpected := range unexpectResolverMaps {
						if _, err := snap.ResolverMaps.List().Find(unexpected.Metadata.Ref().Strings()); err == nil {
							continue drain
						}
					}
					break drain
				case err := <-errs:
					Expect(err).NotTo(HaveOccurred())
				case <-time.After(time.Second * 10):
					nsList1, _ := resolverMapClient.List(namespace1, clients.ListOpts{})
					nsList2, _ := resolverMapClient.List(namespace2, clients.ListOpts{})
					combined := nsList1.ByNamespace()
					combined.Add(nsList2...)
					Fail("expected final snapshot before 10 seconds. expected " + log.Sprintf("%v", combined))
				}
			}
		}
		resolverMap1a, err := resolverMapClient.Write(NewResolverMap(namespace1, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		resolverMap1b, err := resolverMapClient.Write(NewResolverMap(namespace2, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotResolverMaps(ResolverMapList{resolverMap1a, resolverMap1b}, nil)
		resolverMap2a, err := resolverMapClient.Write(NewResolverMap(namespace1, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		resolverMap2b, err := resolverMapClient.Write(NewResolverMap(namespace2, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotResolverMaps(ResolverMapList{resolverMap1a, resolverMap1b, resolverMap2a, resolverMap2b}, nil)

		err = resolverMapClient.Delete(resolverMap2a.Metadata.Namespace, resolverMap2a.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = resolverMapClient.Delete(resolverMap2b.Metadata.Namespace, resolverMap2b.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotResolverMaps(ResolverMapList{resolverMap1a, resolverMap1b}, ResolverMapList{resolverMap2a, resolverMap2b})

		err = resolverMapClient.Delete(resolverMap1a.Metadata.Namespace, resolverMap1a.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = resolverMapClient.Delete(resolverMap1b.Metadata.Namespace, resolverMap1b.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotResolverMaps(nil, ResolverMapList{resolverMap1a, resolverMap1b, resolverMap2a, resolverMap2b})

		/*
			Schema
		*/

		assertSnapshotSchemas := func(expectSchemas SchemaList, unexpectSchemas SchemaList) {
		drain:
			for {
				select {
				case snap = <-snapshots:
					for _, expected := range expectSchemas {
						if _, err := snap.Schemas.List().Find(expected.Metadata.Ref().Strings()); err != nil {
							continue drain
						}
					}
					for _, unexpected := range unexpectSchemas {
						if _, err := snap.Schemas.List().Find(unexpected.Metadata.Ref().Strings()); err == nil {
							continue drain
						}
					}
					break drain
				case err := <-errs:
					Expect(err).NotTo(HaveOccurred())
				case <-time.After(time.Second * 10):
					nsList1, _ := schemaClient.List(namespace1, clients.ListOpts{})
					nsList2, _ := schemaClient.List(namespace2, clients.ListOpts{})
					combined := nsList1.ByNamespace()
					combined.Add(nsList2...)
					Fail("expected final snapshot before 10 seconds. expected " + log.Sprintf("%v", combined))
				}
			}
		}
		schema1a, err := schemaClient.Write(NewSchema(namespace1, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		schema1b, err := schemaClient.Write(NewSchema(namespace2, name1), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotSchemas(SchemaList{schema1a, schema1b}, nil)
		schema2a, err := schemaClient.Write(NewSchema(namespace1, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		schema2b, err := schemaClient.Write(NewSchema(namespace2, name2), clients.WriteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotSchemas(SchemaList{schema1a, schema1b, schema2a, schema2b}, nil)

		err = schemaClient.Delete(schema2a.Metadata.Namespace, schema2a.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = schemaClient.Delete(schema2b.Metadata.Namespace, schema2b.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotSchemas(SchemaList{schema1a, schema1b}, SchemaList{schema2a, schema2b})

		err = schemaClient.Delete(schema1a.Metadata.Namespace, schema1a.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())
		err = schemaClient.Delete(schema1b.Metadata.Namespace, schema1b.Metadata.Name, clients.DeleteOpts{Ctx: ctx})
		Expect(err).NotTo(HaveOccurred())

		assertSnapshotSchemas(nil, SchemaList{schema1a, schema1b, schema2a, schema2b})
	})
})
