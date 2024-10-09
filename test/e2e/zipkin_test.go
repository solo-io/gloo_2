package e2e_test

import (
	"context"
	"fmt"
	"html"
	"io"
	"net/http"
	"time"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/solo-io/gloo/test/services/envoy"
	"go.opentelemetry.io/collector/pdata/ptrace/ptraceotlp"

	"github.com/solo-io/gloo/test/testutils"

	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/gloosnapshot"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/hcm"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/transformation"
	"github.com/solo-io/gloo/projects/gloo/pkg/translator"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"

	gatewayv1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1"
	gatewaydefaults "github.com/solo-io/gloo/projects/gateway/pkg/defaults"
	envoytrace_gloo "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/config/trace/v3"
	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	static_plugin_gloo "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/static"
	"github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/tracing"
	gloohelpers "github.com/solo-io/gloo/test/helpers"
	"github.com/solo-io/gloo/test/v1helpers"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/solo-io/gloo/test/gomega"

	"github.com/solo-io/gloo/test/services"
)

const (
	tracingCollectorPort         = 9411
	tracingCollectorUpstreamName = "tracing-collector"
	openTelemetryCollectionPath  = "/opentelemetry.proto.collector.trace.v1.TraceService/Export"
	zipkinCollectionPath         = "/api/v2/spans"
)

var _ = Describe("Tracing config loading", Serial, func() {

	// These tests use the Serial decorator because they rely on a hard-coded port for the tracing collector (9411)

	var (
		ctx           context.Context
		cancel        context.CancelFunc
		envoyInstance *envoy.Instance
	)

	BeforeEach(func() {
		ctx, cancel = context.WithCancel(context.Background())

		envoyInstance = envoyFactory.NewInstance()
	})

	AfterEach(func() {
		cancel()
	})

	Context("Tracing defined on Envoy bootstrap", func() {

		BeforeEach(func() {
			testutils.ValidateRequirementsAndNotifyGinkgo(
				testutils.LinuxOnly("Uses 127.0.0.1"),
			)
		})

		It("should send trace msgs to the zipkin server", func() {
			err := envoyInstance.RunWithConfigFile(ctx, int(envoyInstance.HttpPort), "./envoyconfigs/zipkin-envoy-conf.yaml")
			Expect(err).NotTo(HaveOccurred())

			// Start a dummy server listening on 9411 for Zipkin requests
			apiHit := make(chan bool, 1)
			zipkinHandler := http.NewServeMux()
			zipkinHandler.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				Expect(r.URL.Path).To(Equal(zipkinCollectionPath))
				fmt.Fprintf(w, "Dummy Zipkin Collector received request on - %q", html.EscapeString(r.URL.Path))
				apiHit <- true
			}))
			startCancellableTracingServer(ctx, fmt.Sprintf("%s:%d", envoyInstance.LocalAddr(), tracingCollectorPort), zipkinHandler)

			// Execute a request against the admin endpoint, as this should result in a trace
			testRequest := createRequestWithTracingEnabled("127.0.0.1", 11082)
			Eventually(func(g Gomega) {
				res, err := testRequest()
				g.Expect(err).NotTo(HaveOccurred())
				g.Expect(res).To(ContainSubstring(`<title>Envoy Admin</title>`))

				g.Eventually(apiHit, 1*time.Second).Should(Receive(BeTrue()))
			}, "10s", ".5s").Should(Succeed(), "Admin endpoint request should result in trace")

		})

	})

	Context("Tracing defined on Gloo resources", func() {

		var (
			testClients  services.TestClients
			testUpstream *v1helpers.TestUpstream
			virtualServiceName string = "vs-test"

			resourcesToCreate *gloosnapshot.ApiSnapshot
		)

		BeforeEach(func() {
			// run gloo
			ro := &services.RunOptions{
				NsToWrite: writeNamespace,
				NsToWatch: []string{"default", writeNamespace},
				WhatToRun: services.What{
					DisableFds: true,
					DisableUds: true,
				},
			}
			testClients = services.RunGlooGatewayUdsFds(ctx, ro)

			// run envoy
			err := envoyInstance.RunWithRole(writeNamespace+"~"+gatewaydefaults.GatewayProxyName, testClients.GlooPort)
			Expect(err).NotTo(HaveOccurred())

			// this is the upstream that will handle requests
			testUpstream = v1helpers.NewTestHttpUpstream(ctx, envoyInstance.LocalAddr())

			vsToTestUpstream := gloohelpers.NewVirtualServiceBuilder().
				WithName(virtualServiceName).
				WithNamespace(writeNamespace).
				WithDomain("test.com").
				WithRoutePrefixMatcher("test", "/").
				WithRouteActionToUpstream("test", testUpstream.Upstream).
				Build()

			// create tracing collector upstream
			tracingCollectorUs := &gloov1.Upstream{
				Metadata: &core.Metadata{
					Name:      tracingCollectorUpstreamName,
					Namespace: writeNamespace,
				},
				UpstreamType: &gloov1.Upstream_Static{
					Static: &static_plugin_gloo.UpstreamSpec{
						Hosts: []*static_plugin_gloo.Host{
							{
								Addr: envoyInstance.LocalAddr(),
								Port: tracingCollectorPort,
							},
						},
					},
				},
			}

			// The set of resources that these tests will generate
			resourcesToCreate = &gloosnapshot.ApiSnapshot{
				Gateways: gatewayv1.GatewayList{
					gatewaydefaults.DefaultGateway(writeNamespace),
				},
				VirtualServices: gatewayv1.VirtualServiceList{
					vsToTestUpstream,
				},
				Upstreams: gloov1.UpstreamList{
					tracingCollectorUs,
					testUpstream.Upstream,
				},
			}
		})

		AfterEach(func() {
			envoyInstance.Clean()
			cancel()
		})

		// Create a handler that writes a boolean to collectorApiChannel when
		// the endpoint is hit
		apiHitHandler := func(collectorApiChannel chan bool, collectionURLPath string) *http.ServeMux {
			tracingCollectorHandler := http.NewServeMux()
			tracingCollectorHandler.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				Expect(r.URL.Path).To(Equal(collectionURLPath))
				fmt.Fprintf(w, "Dummy tracing Collector received request on - %q", html.EscapeString(r.URL.Path))
				collectorApiChannel <- true
			}))
			return tracingCollectorHandler
		}

		otlpHandler := func(exportedTracesChannel chan *ptraceotlp.ExportRequest) *http.ServeMux {
			handler := http.NewServeMux()
			handler.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				fmt.Printf("Got a request! %#v\n", r) // TODO DELETE ME

				reqBody, err := io.ReadAll(r.Body)
				ExpectWithOffset(1, err).NotTo(HaveOccurred())

				exportedTrace := ptraceotlp.NewExportRequest()
				err = exportedTrace.UnmarshalProto(reqBody)
				ExpectWithOffset(1, err).NotTo(HaveOccurred())
				exportedTracesChannel <- &exportedTrace
			}))
			return handler
		}

		startTracingCollectionServer := func(handler *http.ServeMux) {
			// Start a dummy server listening on 9411 for tracing requests
			startCancellableTracingServerHttp2(
				ctx, fmt.Sprintf("%s:%d", envoyInstance.LocalAddr(), tracingCollectorPort), handler)

			// Create Resources
			err := testClients.WriteSnapshot(ctx, resourcesToCreate)
			Expect(err).NotTo(HaveOccurred())

			// Wait for a proxy to be accepted
			gloohelpers.EventuallyResourceAccepted(func() (resources.InputResource, error) {
				return testClients.ProxyClient.Read(writeNamespace, gatewaydefaults.GatewayProxyName, clients.ReadOpts{Ctx: ctx})
			})

			// Ensure the testUpstream is reachable
			v1helpers.ExpectCurlWithOffset(
				0,
				v1helpers.CurlRequest{
					RootCA: nil,
					Port:   envoyInstance.HttpPort,
					Host:   "test.com", // to match the vs-test
					Path:   "/",
					Body:   []byte("solo.io test"),
				},
				v1helpers.CurlResponse{
					Status:  http.StatusOK,
					Message: "solo.io test",
				},
			)
		}

		It("should send trace msgs with valid opentelemetry provider (collector_ref)", func() {
			collectorApiHit := make(chan bool, 1)
			startTracingCollectionServer(apiHitHandler(collectorApiHit, openTelemetryCollectionPath))

			err := gloohelpers.PatchResource(
				ctx,
				&core.ResourceRef{
					Name:      gatewaydefaults.GatewayProxyName,
					Namespace: writeNamespace,
				},
				func(resource resources.Resource) resources.Resource {
					gw := resource.(*gatewayv1.Gateway)
					gw.GetHttpGateway().Options = &gloov1.HttpListenerOptions{
						HttpConnectionManagerSettings: &hcm.HttpConnectionManagerSettings{
							Tracing: &tracing.ListenerTracingSettings{
								ProviderConfig: &tracing.ListenerTracingSettings_OpenTelemetryConfig{
									OpenTelemetryConfig: &envoytrace_gloo.OpenTelemetryConfig{
										CollectorCluster: &envoytrace_gloo.OpenTelemetryConfig_CollectorUpstreamRef{
											CollectorUpstreamRef: &core.ResourceRef{
												Name:      tracingCollectorUpstreamName,
												Namespace: writeNamespace,
											},
										},
									},
								},
							},
						},
					}
					return gw
				},
				testClients.GatewayClient.BaseClient(),
			)
			Expect(err).NotTo(HaveOccurred())

			testRequest := createRequestWithTracingEnabled("localhost", envoyInstance.HttpPort)
			Eventually(func(g Gomega) {
				g.Eventually(testRequest, DefaultEventuallyTimeout, DefaultEventuallyPollingInterval).Should(BeEmpty())
				g.Eventually(collectorApiHit, DefaultEventuallyTimeout, DefaultEventuallyPollingInterval).Should(Receive())
			}, time.Second*10, time.Second).Should(Succeed(), "tracing server should receive trace request")
		})

		FIt("should modify span names when span name transformer is configured", func() {
			exportedTraces := make(chan *ptraceotlp.ExportRequest, 1)
			startTracingCollectionServer(otlpHandler(exportedTraces))

			err := gloohelpers.PatchResource(
				ctx,
				&core.ResourceRef{
					Name: virtualServiceName,
					Namespace: writeNamespace,
				},
				func(resource resources.Resource) resources.Resource {
					vs := resource.(*gatewayv1.VirtualService)
					vs.VirtualHost.Options = &gloov1.VirtualHostOptions{
						StagedTransformations: &transformation.TransformationStages{
							Regular: &transformation.RequestResponseTransformations{
								RequestTransforms: []*transformation.RequestMatch{{
									RequestTransformation: &transformation.Transformation{
										TransformationType: &transformation.Transformation_TransformationTemplate{
											TransformationTemplate: &transformation.TransformationTemplate{
												SpanTransformer: &transformation.TransformationTemplate_SpanTransformer{
													Name: &transformation.InjaTemplate{
														Text: "TESTSPANNAME123123", // FIXME `{{ header("Host") }}`,
													},
												},
											},
										},
									},
								}},
							},
						},
					}
					return vs
				},
				testClients.VirtualServiceClient.BaseClient(),
			)
			Expect(err).NotTo(HaveOccurred())

			Eventually(func(g Gomega) {
				cfg, err := envoyInstance.ConfigDump()
				g.Expect(err).NotTo(HaveOccurred())
				g.Expect(cfg).To(ContainSubstring("TESTSPANNAME123123"))
			}, "12s", "2s")

			err = gloohelpers.PatchResource(
				ctx,
				&core.ResourceRef{
					Name:      gatewaydefaults.GatewayProxyName,
					Namespace: writeNamespace,
				},
				func(resource resources.Resource) resources.Resource {
					gw := resource.(*gatewayv1.Gateway)
					gw.GetHttpGateway().Options = &gloov1.HttpListenerOptions{
						HttpConnectionManagerSettings: &hcm.HttpConnectionManagerSettings{
							Tracing: &tracing.ListenerTracingSettings{
								ProviderConfig: &tracing.ListenerTracingSettings_OpenTelemetryConfig{
									OpenTelemetryConfig: &envoytrace_gloo.OpenTelemetryConfig{
										CollectorCluster: &envoytrace_gloo.OpenTelemetryConfig_CollectorUpstreamRef{
											CollectorUpstreamRef: &core.ResourceRef{
												Name:      tracingCollectorUpstreamName,
												Namespace: writeNamespace,
											},
										},
									},
								},
							},
						},
					}
					return gw
				},
				testClients.GatewayClient.BaseClient(),
			)
			Expect(err).NotTo(HaveOccurred())

			testRequest := createRequestWithTracingEnabled("localhost", envoyInstance.HttpPort)
			Eventually(func(g Gomega) {
				g.Eventually(testRequest, DefaultEventuallyTimeout, DefaultEventuallyPollingInterval).Should(BeEmpty())
				var exportedTrace *ptraceotlp.ExportRequest
				g.Eventually(exportedTraces, DefaultEventuallyTimeout, DefaultEventuallyPollingInterval).Should(Receive(&exportedTrace))
				for exportedTrace := range exportedTraces {
					resourceSpans := exportedTrace.Traces().ResourceSpans()
					for resourceSpanIter := 0; resourceSpanIter < resourceSpans.Len(); resourceSpanIter++ {
						resourceSpan := resourceSpans.At(resourceSpanIter)
						scopeSpans := resourceSpan.ScopeSpans()
						for scopeSpanIter := 0; scopeSpanIter < scopeSpans.Len(); scopeSpanIter++ {
							scopeSpan := scopeSpans.At(scopeSpanIter)
							spans := scopeSpan.Spans()
							for spanIter := 0; spanIter < spans.Len(); spanIter++ {
								span := spans.At(spanIter)
								if span.Name() == "TESTSPANNAME123123" {
									return
								}
							}
						}
					}
				}
			}, time.Second*10, time.Second).Should(Succeed(), "tracing server should receive trace request")
		})

		It("should send trace msgs with valid opentelemetry provider (cluster_name)", func() {
			collectorApiHit := make(chan bool, 1)
			startTracingCollectionServer(apiHitHandler(collectorApiHit, openTelemetryCollectionPath))

			err := gloohelpers.PatchResource(
				ctx,
				&core.ResourceRef{
					Name:      gatewaydefaults.GatewayProxyName,
					Namespace: writeNamespace,
				},
				func(resource resources.Resource) resources.Resource {
					gw := resource.(*gatewayv1.Gateway)
					gw.GetHttpGateway().Options = &gloov1.HttpListenerOptions{
						HttpConnectionManagerSettings: &hcm.HttpConnectionManagerSettings{
							Tracing: &tracing.ListenerTracingSettings{
								ProviderConfig: &tracing.ListenerTracingSettings_OpenTelemetryConfig{
									OpenTelemetryConfig: &envoytrace_gloo.OpenTelemetryConfig{
										CollectorCluster: &envoytrace_gloo.OpenTelemetryConfig_ClusterName{
											ClusterName: translator.UpstreamToClusterName(&core.ResourceRef{
												Name:      tracingCollectorUpstreamName,
												Namespace: writeNamespace,
											}),
										},
									},
								},
							},
						},
					}
					return gw
				},
				testClients.GatewayClient.BaseClient(),
			)
			Expect(err).NotTo(HaveOccurred())

			testRequest := createRequestWithTracingEnabled("localhost", envoyInstance.HttpPort)
			Eventually(func(g Gomega) {
				g.Eventually(testRequest, DefaultEventuallyTimeout, DefaultEventuallyPollingInterval).Should(BeEmpty())
				g.Eventually(collectorApiHit, DefaultEventuallyTimeout, DefaultEventuallyPollingInterval).Should(Receive())
			}, time.Second*10, time.Second).Should(Succeed(), "tracing server should receive trace request")
		})

		It("should not send trace msgs with nil provider", func() {
			collectorApiHit := make(chan bool, 1)
			startTracingCollectionServer(apiHitHandler(collectorApiHit, zipkinCollectionPath))

			err := gloohelpers.PatchResource(
				ctx,
				&core.ResourceRef{
					Name:      gatewaydefaults.GatewayProxyName,
					Namespace: writeNamespace,
				},
				func(resource resources.Resource) resources.Resource {
					gw := resource.(*gatewayv1.Gateway)
					gw.GetHttpGateway().Options = &gloov1.HttpListenerOptions{
						HttpConnectionManagerSettings: &hcm.HttpConnectionManagerSettings{
							Tracing: nil,
						},
					}
					return gw
				},
				testClients.GatewayClient.BaseClient(),
			)
			Expect(err).NotTo(HaveOccurred())

			testRequest := createRequestWithTracingEnabled("localhost", envoyInstance.HttpPort)
			Eventually(func(g Gomega) {
				g.Eventually(testRequest, DefaultEventuallyTimeout, DefaultEventuallyPollingInterval).Should(BeEmpty())
				g.Eventually(collectorApiHit, DefaultEventuallyTimeout, DefaultEventuallyPollingInterval).Should(Not(Receive()))
			}, time.Second*5, time.Millisecond*250).Should(Succeed(), "zipkin server should not receive trace request")
		})

		It("should send trace msgs with valid zipkin provider (collector_ref)", func() {
			collectorApiHit := make(chan bool, 1)
			startTracingCollectionServer(apiHitHandler(collectorApiHit, zipkinCollectionPath))

			err := gloohelpers.PatchResource(
				ctx,
				&core.ResourceRef{
					Name:      gatewaydefaults.GatewayProxyName,
					Namespace: writeNamespace,
				},
				func(resource resources.Resource) resources.Resource {
					gw := resource.(*gatewayv1.Gateway)
					gw.GetHttpGateway().Options = &gloov1.HttpListenerOptions{
						HttpConnectionManagerSettings: &hcm.HttpConnectionManagerSettings{
							Tracing: &tracing.ListenerTracingSettings{
								ProviderConfig: &tracing.ListenerTracingSettings_ZipkinConfig{
									ZipkinConfig: &envoytrace_gloo.ZipkinConfig{
										CollectorCluster: &envoytrace_gloo.ZipkinConfig_CollectorUpstreamRef{
											CollectorUpstreamRef: &core.ResourceRef{
												Name:      tracingCollectorUpstreamName,
												Namespace: writeNamespace,
											},
										},
										CollectorEndpoint:        zipkinCollectionPath,
										CollectorEndpointVersion: envoytrace_gloo.ZipkinConfig_HTTP_JSON,
									},
								},
							},
						},
					}
					return gw
				},
				testClients.GatewayClient.BaseClient(),
			)
			Expect(err).NotTo(HaveOccurred())

			testRequest := createRequestWithTracingEnabled("localhost", envoyInstance.HttpPort)
			Eventually(func(g Gomega) {
				g.Eventually(testRequest, DefaultEventuallyTimeout, DefaultEventuallyPollingInterval).Should(BeEmpty())
				g.Eventually(collectorApiHit, DefaultEventuallyTimeout, DefaultEventuallyPollingInterval).Should(Receive())
			}, time.Second*10, time.Second).Should(Succeed(), "tracing server should receive trace request")
		})

		It("should send trace msgs with valid zipkin provider (cluster_name)", func() {
			collectorApiHit := make(chan bool, 1)
			startTracingCollectionServer(apiHitHandler(collectorApiHit, zipkinCollectionPath))

			err := gloohelpers.PatchResource(
				ctx,
				&core.ResourceRef{
					Name:      gatewaydefaults.GatewayProxyName,
					Namespace: writeNamespace,
				},
				func(resource resources.Resource) resources.Resource {
					gw := resource.(*gatewayv1.Gateway)
					gw.GetHttpGateway().Options = &gloov1.HttpListenerOptions{
						HttpConnectionManagerSettings: &hcm.HttpConnectionManagerSettings{
							Tracing: &tracing.ListenerTracingSettings{
								ProviderConfig: &tracing.ListenerTracingSettings_ZipkinConfig{
									ZipkinConfig: &envoytrace_gloo.ZipkinConfig{
										CollectorCluster: &envoytrace_gloo.ZipkinConfig_ClusterName{
											ClusterName: translator.UpstreamToClusterName(&core.ResourceRef{
												Name:      tracingCollectorUpstreamName,
												Namespace: writeNamespace,
											}),
										},
										CollectorEndpoint:        zipkinCollectionPath,
										CollectorEndpointVersion: envoytrace_gloo.ZipkinConfig_HTTP_JSON,
									},
								},
							},
						},
					}
					return gw
				},
				testClients.GatewayClient.BaseClient(),
			)
			Expect(err).NotTo(HaveOccurred())

			testRequest := createRequestWithTracingEnabled("localhost", envoyInstance.HttpPort)
			Eventually(func(g Gomega) {
				g.Eventually(testRequest, DefaultEventuallyTimeout, DefaultEventuallyPollingInterval).Should(BeEmpty())
				g.Eventually(collectorApiHit, DefaultEventuallyTimeout, DefaultEventuallyPollingInterval).Should(Receive())
			}, time.Second*10, time.Second).Should(Succeed(), "zipkin server should receive trace request")
		})

		It("should error with invalid zipkin provider", func() {
			collectorApiHit := make(chan bool, 1)
			startTracingCollectionServer(apiHitHandler(collectorApiHit, zipkinCollectionPath))

			err := gloohelpers.PatchResource(
				ctx,
				&core.ResourceRef{
					Name:      gatewaydefaults.GatewayProxyName,
					Namespace: writeNamespace,
				},
				func(resource resources.Resource) resources.Resource {
					gw := resource.(*gatewayv1.Gateway)
					gw.GetHttpGateway().Options = &gloov1.HttpListenerOptions{
						HttpConnectionManagerSettings: &hcm.HttpConnectionManagerSettings{
							Tracing: &tracing.ListenerTracingSettings{
								ProviderConfig: &tracing.ListenerTracingSettings_ZipkinConfig{
									ZipkinConfig: &envoytrace_gloo.ZipkinConfig{
										CollectorCluster: &envoytrace_gloo.ZipkinConfig_CollectorUpstreamRef{
											CollectorUpstreamRef: nil,
										},
										CollectorEndpoint:        zipkinCollectionPath,
										CollectorEndpointVersion: envoytrace_gloo.ZipkinConfig_HTTP_JSON,
									},
								},
							},
						},
					}
					return gw
				},
				testClients.GatewayClient.BaseClient(),
			)
			Expect(err).NotTo(HaveOccurred())

			// ensure the proxy is never updated with the invalid configuration
			Consistently(func(g Gomega) int {
				tracingConfigsFound := 0

				proxy, err := testClients.ProxyClient.Read(writeNamespace, gatewaydefaults.GatewayProxyName, clients.ReadOpts{Ctx: ctx})
				g.Expect(err).NotTo(HaveOccurred())

				for _, l := range proxy.GetListeners() {
					if l.GetHttpListener().GetOptions().GetHttpConnectionManagerSettings().GetTracing() != nil {
						tracingConfigsFound += 1
					}
				}
				return tracingConfigsFound
			}, time.Second*3, time.Second).Should(Equal(0))
		})
	})

})

func startCancellableTracingServerHttp2(serverContext context.Context, address string, handler http.Handler) {
	tracingServer := http.Server{}
	http2Server := http2.Server{}
	err := http2.ConfigureServer(&tracingServer, &http2Server)
	ExpectWithOffset(1, err).NotTo(HaveOccurred())

	// Start a goroutine to handle requests
	go func() {
		defer GinkgoRecover()
		h2cHandler := h2c.NewHandler(handler, &http2.Server{})
		if err := http.ListenAndServe(address, h2cHandler); err != nil && err != http.ErrServerClosed {
			ExpectWithOffset(1, err).NotTo(HaveOccurred())
		}
	}()

	// Start a goroutine to shutdown the server
	go func(serverCtx context.Context) {
		defer GinkgoRecover()

		<-serverCtx.Done()
		// tracingServer.Shutdown hangs with opentelemetry tests, probably
		// because the agent leaves the connection open. There's no need for a
		// graceful shutdown anyway, so just force it using Close() instead
		tracingServer.Close()
	}(serverContext)
}

func startCancellableTracingServer(serverContext context.Context, address string, handler http.Handler) {
	tracingServer := &http.Server{
		Addr:    address,
		Handler: handler,
	}

	// Start a goroutine to handle requests
	go func() {
		defer GinkgoRecover()
		if err := tracingServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			ExpectWithOffset(1, err).NotTo(HaveOccurred())
		}
	}()

	// Start a goroutine to shutdown the server
	go func(serverCtx context.Context) {
		defer GinkgoRecover()

		<-serverCtx.Done()
		// tracingServer.Shutdown hangs with opentelemetry tests, probably
		// because the agent leaves the connection open. There's no need for a
		// graceful shutdown anyway, so just force it using Close() instead
		tracingServer.Close()
	}(serverContext)
}

func createRequestWithTracingEnabled(address string, port uint32) func() (string, error) {
	return func() (string, error) {
		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://%s:%d/", address, port), nil)
		if err != nil {
			return "", err
		}
		req.Header.Set("Content-Type", "application/json")

		// Set a random trace ID
		req.Header.Set("x-client-trace-id", "test-trace-id-1234567890")

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return "", err
		}
		defer res.Body.Close()
		body, err := io.ReadAll(res.Body)
		return string(body), err
	}
}
