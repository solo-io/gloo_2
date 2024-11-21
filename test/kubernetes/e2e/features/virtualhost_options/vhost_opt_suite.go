package virtualhost_options

import (
	"context"
	"net/http"

	"github.com/onsi/gomega"
	"github.com/onsi/gomega/gstruct"
	"github.com/stretchr/testify/suite"

	"github.com/solo-io/gloo/pkg/utils/kubeutils"
	"github.com/solo-io/gloo/pkg/utils/requestutils/curl"
	"github.com/solo-io/gloo/projects/gloo/pkg/defaults"
	"github.com/solo-io/gloo/test/gomega/matchers"
	"github.com/solo-io/gloo/test/helpers"
	"github.com/solo-io/gloo/test/kubernetes/e2e"
	testdefaults "github.com/solo-io/gloo/test/kubernetes/e2e/defaults"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var _ e2e.NewSuiteFunc = NewTestingSuite

// testingSuite is the entire Suite of tests for the "VirtualHostOptions" feature
type testingSuite struct {
	suite.Suite

	ctx context.Context

	// testInstallation contains all the metadata/utilities necessary to execute a series of tests
	// against an installation of Gloo Gateway
	testInstallation *e2e.TestInstallation
}

func NewTestingSuite(ctx context.Context, testInst *e2e.TestInstallation) suite.TestingSuite {
	return &testingSuite{
		ctx:              ctx,
		testInstallation: testInst,
	}
}

func (s *testingSuite) SetupSuite() {
	// Check that the common setup manifest is applied
	for _, manifest := range setupManifests {
		err := s.testInstallation.Actions.Kubectl().ApplyFile(s.ctx, manifest)
		s.NoError(err, "can apply "+manifest)
	}

	s.testInstallation.Assertions.EventuallyObjectsExist(s.ctx, proxyService, proxyDeployment,
		exampleSvc, nginxPod, testdefaults.CurlPod)

	// Check that test resources are running
	s.testInstallation.Assertions.EventuallyPodsRunning(s.ctx, nginxPod.ObjectMeta.GetNamespace(),
		metav1.ListOptions{
			LabelSelector: "app.kubernetes.io/name=nginx",
		})
	s.testInstallation.Assertions.EventuallyPodsRunning(s.ctx, proxyDeployment.ObjectMeta.GetNamespace(),
		metav1.ListOptions{
			LabelSelector: "app.kubernetes.io/name=gloo-proxy-gw",
		})
	s.testInstallation.Assertions.EventuallyPodsRunning(s.ctx, testdefaults.CurlPod.GetNamespace(),
		metav1.ListOptions{
			LabelSelector: "app.kubernetes.io/name=curl",
		})
}

func (s *testingSuite) TearDownSuite() {
	// Check that the common setup manifest is deleted
	for _, manifest := range setupManifests {
		output, err := s.testInstallation.Actions.Kubectl().DeleteFileWithOutput(s.ctx, manifest)
		s.NoError(err, "can delete "+manifest)
		s.testInstallation.Assertions.ExpectObjectDeleted(manifest, err, output)
	}
}

// TestConfigureVirtualHostOptions tests the basic functionality of VirtualHostOptions using a single VHO
func (s *testingSuite) TestConfigureVirtualHostOptions() {
	s.T().Cleanup(func() {
		output, err := s.testInstallation.Actions.Kubectl().DeleteFileWithOutput(s.ctx, manifestVhoRemoveXBar)
		s.testInstallation.Assertions.ExpectObjectDeleted(manifestVhoRemoveXBar, err, output)
	})

	err := s.testInstallation.Actions.Kubectl().ApplyFile(s.ctx, manifestVhoRemoveXBar)
	s.NoError(err, "can apply "+manifestVhoRemoveXBar)

	// Check status is accepted on VirtualHostOption
	s.testInstallation.Assertions.EventuallyResourceStatusMatchesState(
		s.getterForMeta(&vhoRemoveXBar),
		core.Status_Accepted,
		defaults.KubeGatewayReporter,
	)

	// Check healthy response with no x-bar header
	s.testInstallation.Assertions.AssertEventualCurlResponse(
		s.ctx,
		testdefaults.CurlPodExecOpt,
		[]curl.Option{
			curl.WithHost(kubeutils.ServiceFQDN(proxyService.ObjectMeta)),
			curl.WithHostHeader("example.com"),
		},
		expectedResponseWithoutXBar)
}

// TestConfigureInvalidVirtualHostOptions confirms that an invalid VirtualHostOption is rejected
func (s *testingSuite) TestConfigureInvalidVirtualHostOptions() {
	s.T().Cleanup(func() {
		output, err := s.testInstallation.Actions.Kubectl().DeleteFileWithOutput(s.ctx, manifestVhoWebhookReject)
		s.testInstallation.Assertions.ExpectObjectDeleted(manifestVhoWebhookReject, err, output)
	})

	err := s.testInstallation.Actions.Kubectl().ApplyFile(s.ctx, manifestVhoWebhookReject)
	s.NoError(err, "can apply "+manifestVhoWebhookReject)

	if !s.testInstallation.Metadata.ValidationAlwaysAccept {
		s.testInstallation.Assertions.ExpectGlooObjectNotExist(
			s.ctx,
			s.getterForMeta(&vhoWebhookReject),
			&vhoWebhookReject,
		)
	} else {
		// Check status is rejected on bad VirtualHostOption
		s.testInstallation.Assertions.EventuallyResourceStatusMatchesState(
			s.getterForMeta(&vhoWebhookReject),
			core.Status_Rejected,
			defaults.KubeGatewayReporter,
		)
	}
}

// TestConfigureVirtualHostOptionsWithSectionName tests a complex scenario where multiple VirtualHostOptions conflicting
// across multiple listeners are applied to a single gateway
//
// The goal here is to test the behavior when multiple VHOs target a gateway with multiple listeners and only some
// conflict. This will generate a warning on the conflicted resource, but the VHO should be attached properly and
// options propagated for the listener.
func (s *testingSuite) TestConfigureVirtualHostOptionsWithSectionNameManualSetup() {
	s.T().Cleanup(func() {
		output, err := s.testInstallation.Actions.Kubectl().DeleteFileWithOutput(s.ctx, manifestVhoRemoveXBar)
		s.testInstallation.Assertions.ExpectObjectDeleted(manifestVhoRemoveXBar, err, output)

		output, err = s.testInstallation.Actions.Kubectl().DeleteFileWithOutput(s.ctx, manifestVHORemoveXBaz)
		s.testInstallation.Assertions.ExpectObjectDeleted(manifestVHORemoveXBaz, err, output)

		output, err = s.testInstallation.Actions.Kubectl().DeleteFileWithOutput(s.ctx, manifestVhoSectionAddXFoo)
		s.testInstallation.Assertions.ExpectObjectDeleted(manifestVhoSectionAddXFoo, err, output)
	})

	// Apply our manifests so we can assert that basic vho exists before applying conflicting VHOs.
	// This is needed because our solo-kit clients currently do not return creationTimestamp
	err := s.testInstallation.Actions.Kubectl().ApplyFile(s.ctx, manifestVhoRemoveXBar)
	s.NoError(err, "can apply "+manifestVhoRemoveXBar)

	// Check status is accepted before moving on to apply conflicting vho
	s.testInstallation.Assertions.EventuallyResourceStatusMatchesState(
		s.getterForMeta(&vhoRemoveXBar),
		core.Status_Accepted,
		defaults.KubeGatewayReporter,
	)

	err = s.testInstallation.Actions.Kubectl().ApplyFile(s.ctx, manifestVHORemoveXBaz)
	s.NoError(err, "can apply "+manifestVHORemoveXBaz)

	err = s.testInstallation.Actions.Kubectl().ApplyFile(s.ctx, manifestVhoSectionAddXFoo)
	s.NoError(err, "can apply "+manifestVhoSectionAddXFoo)

	// Check status is accepted on VirtualHostOption with section name
	s.testInstallation.Assertions.EventuallyResourceStatusMatchesState(
		s.getterForMeta(&vhoSectionAddXFoo),
		core.Status_Accepted,
		defaults.KubeGatewayReporter,
	)

	// Check status is warning on VirtualHostOption not selected for attachment
	s.testInstallation.Assertions.EventuallyResourceStatusMatchesWarningReasons(
		s.getterForMeta(&vhoRemoveXBaz),
		[]string{"conflict with more specific or older VirtualHostOptions"},
		defaults.KubeGatewayReporter,
	)

	// Check healthy response with added foo header to listener targeted by sectionName
	s.testInstallation.Assertions.AssertEventualCurlResponse(
		s.ctx,
		testdefaults.CurlPodExecOpt,
		[]curl.Option{
			curl.WithHost(kubeutils.ServiceFQDN(proxyService.ObjectMeta)),
			curl.WithHostHeader("example.com"),
			curl.WithPort(8080),
		},
		expectedResponseWithXFoo)

	// Check status is warning on VirtualHostOption with conflicting attachment,
	// despite being properly attached to another listener
	s.testInstallation.Assertions.EventuallyResourceStatusMatchesWarningReasons(
		s.getterForMeta(&vhoRemoveXBar),
		[]string{"conflict with more specific or older VirtualHostOptions"},
		defaults.KubeGatewayReporter,
	)

	// Check healthy response with x-bar removed to listener NOT targeted by sectionName
	s.testInstallation.Assertions.AssertEventualCurlResponse(
		s.ctx,
		testdefaults.CurlPodExecOpt,
		[]curl.Option{
			curl.WithHost(kubeutils.ServiceFQDN(proxyService.ObjectMeta)),
			curl.WithHostHeader("example.com"),
			curl.WithPort(8081),
		},
		expectedResponseWithoutXBar)
}

// TestMultipleVirtualHostOptionsSetup tests a complex scenario where multiple VirtualHostOptions conflict
//
// The goal here is to test the behavior when multiple VHOs are targeting a gateway without sectionName.
// The expected behavior is that the oldest resource is used
func (s *testingSuite) TestMultipleVirtualHostOptionsSetup() {
	s.T().Cleanup(func() {
		output, err := s.testInstallation.Actions.Kubectl().DeleteFileWithOutput(s.ctx, manifestVhoRemoveXBar)
		s.testInstallation.Assertions.ExpectObjectDeleted(manifestVhoRemoveXBar, err, output)

		output, err = s.testInstallation.Actions.Kubectl().DeleteFileWithOutput(s.ctx, manifestVHORemoveXBaz)
		s.testInstallation.Assertions.ExpectObjectDeleted(manifestVHORemoveXBaz, err, output)
	})

	// Manually apply our manifests so we can assert that basic vho exists before applying extra vho.
	// This is needed because our solo-kit clients currently do not return creationTimestamp
	err := s.testInstallation.Actions.Kubectl().ApplyFile(s.ctx, manifestVhoRemoveXBar)
	s.NoError(err, "can apply "+manifestVhoRemoveXBar)

	err = s.testInstallation.Actions.Kubectl().ApplyFile(s.ctx, manifestVHORemoveXBaz)
	s.NoError(err, "can apply "+manifestVHORemoveXBaz)

	// Check status is warning on newer VirtualHostOption not selected for attachment
	s.testInstallation.Assertions.EventuallyResourceStatusMatchesWarningReasons(
		s.getterForMeta(&vhoRemoveXBaz),
		[]string{"conflict with more specific or older VirtualHostOptions"},
		defaults.KubeGatewayReporter,
	)

	// Check status is accepted on older VirtualHostOption
	s.testInstallation.Assertions.EventuallyResourceStatusMatchesState(
		s.getterForMeta(&vhoRemoveXBar),
		core.Status_Accepted,
		defaults.KubeGatewayReporter,
	)

	// Check healthy response with no x-bar header
	s.testInstallation.Assertions.AssertEventualCurlResponse(
		s.ctx,
		testdefaults.CurlPodExecOpt,
		[]curl.Option{
			curl.WithHost(kubeutils.ServiceFQDN(proxyService.ObjectMeta)),
			curl.WithHostHeader("example.com"),
		},
		expectedResponseWithoutXBar)
}

// TestOptionsMerge tests shallow merging of VirtualHostOptions larger in the precedence chain
func (s *testingSuite) TestOptionsMerge() {
	s.T().Cleanup(func() {
		output, err := s.testInstallation.Actions.Kubectl().DeleteFileWithOutput(s.ctx, manifestVhoRemoveXBar)
		s.testInstallation.Assertions.ExpectObjectDeleted(manifestVhoRemoveXBar, err, output)

		output, err = s.testInstallation.Actions.Kubectl().DeleteFileWithOutput(s.ctx, manifestVhoMergeRemoveXBaz)
		s.testInstallation.Assertions.ExpectObjectDeleted(manifestVhoMergeRemoveXBaz, err, output)
	})

	_, err := s.testInstallation.Actions.Kubectl().ApplyFileWithOutput(s.ctx, manifestVhoRemoveXBar)
	s.Require().NoError(err)

	_, err = s.testInstallation.Actions.Kubectl().ApplyFileWithOutput(s.ctx, manifestVhoMergeRemoveXBaz)
	s.Require().NoError(err)

	s.testInstallation.Assertions.EventuallyResourceStatusMatchesState(
		s.getterForMeta(&vhoMergeRemoveXBaz),
		core.Status_Accepted,
		defaults.KubeGatewayReporter,
	)

	s.testInstallation.Assertions.EventuallyResourceStatusMatchesState(
		s.getterForMeta(&vhoRemoveXBar),
		core.Status_Accepted,
		defaults.KubeGatewayReporter,
	)

	s.testInstallation.Assertions.AssertEventualCurlResponse(
		s.ctx,
		testdefaults.CurlPodExecOpt,
		[]curl.Option{
			curl.WithHost(kubeutils.ServiceFQDN(proxyService.ObjectMeta)),
			curl.WithHostHeader("example.com"),
		},
		// Expect:
		// - x-bar header to be removed by vho-remove-x-bar
		// - x-baz header to not be removed as the option conflicts with vho-remove-x-bar and is not merged
		// - x-envoy-attempt-count header to be added by vho-merge-remove-x-baz.yaml
		&matchers.HttpResponse{
			StatusCode: http.StatusOK,
			Custom: gomega.And(
				gomega.Not(matchers.ContainHeaderKeys([]string{"x-bar"})),
				matchers.ContainHeaderKeys([]string{"x-baz"}),
				matchers.ContainHeaderKeys([]string{"x-envoy-attempt-count"}),
			),
			Body: gstruct.Ignore(),
		},
	)
}

func (s *testingSuite) getterForMeta(meta *metav1.ObjectMeta) helpers.InputResourceGetter {
	return func() (resources.InputResource, error) {
		return s.testInstallation.ResourceClients.VirtualHostOptionClient().Read(meta.GetNamespace(),
			meta.GetName(), clients.ReadOpts{})
	}
}
