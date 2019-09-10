package e2e_test

import (
	"testing"

	"github.com/solo-io/go-utils/testutils"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/solo-io/solo-kit/test/helpers"
	"github.com/solo-io/solo-projects/test/services"

	"github.com/solo-io/gloo/projects/gloo/pkg/defaults"
)

var (
	envoyFactory  *services.EnvoyFactory
	consulFactory *services.ConsulFactory
)

var _ = BeforeSuite(func() {
	var err error
	envoyFactory, err = services.NewEnvoyFactory()
	Expect(err).NotTo(HaveOccurred())
	consulFactory, err = services.NewConsulFactory()
	Expect(err).NotTo(HaveOccurred())

})

var _ = AfterSuite(func() {
	envoyFactory.Clean()
	consulFactory.Clean()
})

func TestE2e(t *testing.T) {

	// set default port to an unprivileged port for local testing.
	defaults.HttpPort = 8081

	testutils.RegisterPreFailHandler(testutils.PrintTrimmedStack)
	helpers.RegisterCommonFailHandlers()
	helpers.SetupLog()
	// RegisterFailHandler(Fail)
	RunSpecs(t, "E2e Suite")
}
