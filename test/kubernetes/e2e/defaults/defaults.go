package defaults

import (
	"context"
	"path/filepath"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/solo-io/gloo/pkg/utils/kubeutils/kubectl"
	"github.com/solo-io/gloo/test/kubernetes/e2e"
	"github.com/solo-io/skv2/codegen/util"
	"github.com/stretchr/testify/assert"
)

type CommonTestSuite interface {
	TestInstallation() *e2e.TestInstallation
	Ctx() context.Context
	Assert() *assert.Assertions
}

var (
	CurlPodExecOpt = kubectl.PodExecOptions{
		Name:      "curl",
		Namespace: "curl",
		Container: "curl",
	}

	CurlPod = &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "curl",
			Namespace: "curl",
		},
	}

	CurlPodManifest = filepath.Join(util.MustGetThisDir(), "testdata", "curl_pod.yaml")

	HttpEchoPod = &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "http-echo",
			Namespace: "http-echo",
		},
	}

	HttpEchoPodManifest = filepath.Join(util.MustGetThisDir(), "testdata", "http_echo.yaml")

	TcpEchoPod = &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "tcp-echo",
			Namespace: "tcp-echo",
		},
	}

	TcpEchoPodManifest = filepath.Join(util.MustGetThisDir(), "testdata", "tcp_echo.yaml")

	NginxPod = &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "nginx",
			Namespace: "nginx",
		},
	}

	NginxSvc = &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "nginx",
			Namespace: "nginx",
		},
	}

	NginxPodManifest = filepath.Join(util.MustGetThisDir(), "testdata", "nginx_pod.yaml")

	NginxResponse = `<!DOCTYPE html>
<html>
<head>
<title>Welcome to nginx!</title>
<style>
html { color-scheme: light dark; }
body { width: 35em; margin: 0 auto;
font-family: Tahoma, Verdana, Arial, sans-serif; }
</style>
</head>
<body>
<h1>Welcome to nginx!</h1>
<p>If you see this page, the nginx web server is successfully installed and
working. Further configuration is required.</p>

<p>For online documentation and support please refer to
<a href="http://nginx.org/">nginx.org</a>.<br/>
Commercial support is available at
<a href="http://nginx.com/">nginx.com</a>.</p>

<p><em>Thank you for using nginx.</em></p>
</body>
</html>`
)

func CurlPodEventuallyRunning(s CommonTestSuite) {
	// Check that test resources are running
	s.TestInstallation().Assertions.EventuallyPodsRunning(s.Ctx(), CurlPod.ObjectMeta.GetNamespace(), metav1.ListOptions{
		LabelSelector: "app.kubernetes.io/name=curl",
	})
}

func InstallCurlPod(s CommonTestSuite) {
	err := s.TestInstallation().Actions.Kubectl().ApplyFile(s.Ctx(), CurlPodManifest)
	s.Assert().NoError(err)
	CurlPodEventuallyRunning(s)
}

func DeleteCurlPod(s CommonTestSuite) {
	output, err := s.TestInstallation().Actions.Kubectl().DeleteFileWithOutput(s.Ctx(), CurlPodManifest)
	s.Assert().NoError(err, "can delete curl pod")
	s.TestInstallation().Assertions.ExpectObjectDeleted(CurlPodManifest, err, output)
}
