package kubernetes

import (
	"context"
	"fmt"
	"reflect"

	"github.com/solo-io/gloo/pkg/utils/envutils"
	"github.com/solo-io/gloo/projects/gloo/constants"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	kubeplugin "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/kubernetes"
	"github.com/solo-io/gloo/projects/gloo/pkg/discovery"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins/kubernetes/serviceconverter"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins/utils"
	"github.com/solo-io/go-utils/contextutils"
	sanitizer "github.com/solo-io/k8s-utils/kubeutils"
	"github.com/solo-io/solo-kit/pkg/errors"
	"github.com/solo-io/solo-kit/pkg/utils/kubeutils"
	corev1 "k8s.io/api/core/v1"
)

type UpstreamConverter interface {
	UpstreamsForService(ctx context.Context, svc *corev1.Service) v1.UpstreamList
}

func DefaultUpstreamConverter() *KubeUpstreamConverter {
	kuc := new(KubeUpstreamConverter)
	kuc.serviceConverters = serviceconverter.DefaultServiceConverters
	return kuc
}

type KubeUpstreamConverter struct {
	serviceConverters []serviceconverter.ServiceConverter
}

// UpstreamsForService is called by the k8s uds plugin to convert a service to list of upstreams
func (uc *KubeUpstreamConverter) UpstreamsForService(ctx context.Context, svc *corev1.Service) v1.UpstreamList {
	var upstreams v1.UpstreamList
	for _, port := range svc.Spec.Ports {
		upstreams = append(upstreams, uc.CreateUpstream(ctx, svc, port))
	}
	return upstreams
}

// CreateUpstream is called by both:
// - discovery (when creating an upstream from a k8s service)
// - controller code that converts services to in-memory upstreams
func (uc *KubeUpstreamConverter) CreateUpstream(ctx context.Context, svc *corev1.Service, port corev1.ServicePort) *v1.Upstream {
	meta := svc.ObjectMeta
	coremeta := kubeutils.FromKubeMeta(meta, false)
	coremeta.ResourceVersion = ""
	coremeta.Name = UpstreamName(meta.Namespace, meta.Name, port.Port, true)
	labels := coremeta.GetLabels()
	coremeta.Labels = make(map[string]string)

	us := &v1.Upstream{
		Metadata: coremeta,
		UpstreamType: &v1.Upstream_Kube{
			Kube: &kubeplugin.UpstreamSpec{
				ServiceName:      meta.Name,
				ServiceNamespace: meta.Namespace,
				ServicePort:      uint32(port.Port),
				Selector:         svc.Spec.Selector,
			},
		},
		DiscoveryMetadata: &v1.DiscoveryMetadata{
			Labels: labels,
		},
	}

	for _, sc := range uc.serviceConverters {
		if err := sc.ConvertService(ctx, svc, port, us); err != nil {
			contextutils.LoggerFrom(ctx).Errorf("error: failed to process service options with err %v", err)
		}
	}

	return us
}

// UpstreamName creates an upstream name from a k8s Service name/namespace/port.
//
// This function is used in the context of both "real" upstreams written to etcd (e.g. upstream creation by UDS)
// as well as "fake" upstreams (i.e. those upstreams we only create in the in-memory input snapshot).
//
// The caller should set `sanitize` to true if this name is being created for a "real" upstream (since
// k8s names/namespaces have some character restrictions), and should set `sanitize` to false if this is
// for an in-memory upstream (so we can make cluster names more parseable by using an invalid k8s character
// as a separator).
//
// The resulting upstream name follows one of two formats:
//  1. <svcNs>_<svcName>_<svcPort> (underscore-separated): only if kubernetes gateway integration is enabled AND
//     `sanitize` is false
//  2. <svcNs>-<svcName>-<svcPort> (hyphen-separated) otherwise.
//
// If kubernetes gateway integration is not enabled (i.e. only edge api is being used), we fall back to the legacy
// (hyphen-separated) format, so as not to cause unexpected changes for existing users.
func UpstreamName(serviceNamespace, serviceName string, servicePort int32, sanitize bool) string {
	if envutils.IsEnvTruthy(constants.GlooGatewayEnableK8sGwControllerEnv) && !sanitize {
		// when kube gateway is enabled, use _ since it makes the service name/namespace easier to
		// parse out later (namespaces/names cannot have _ in them)
		return fmt.Sprintf("%s_%s_%v", serviceNamespace, serviceName, servicePort)
	}
	// for edge api, leave the cluster names as they were (we don't want to cause unexpected changes on upgrade)
	return sanitizer.SanitizeNameV2(fmt.Sprintf("%s-%s-%v", serviceNamespace, serviceName, servicePort))
}

func skip(svc *corev1.Service, opts discovery.Opts) bool {
	// ilackarms: allow user to override the skip with an annotation
	// force discovery for a service with no selector
	if svc.ObjectMeta.Annotations[discoveryAnnotationKey] == discoveryAnnotationTrue {
		return false
	}
	// note: ilackarms: IgnoredServices is not set anywhere
	for _, name := range opts.KubeOpts.IgnoredServices {
		if svc.Name == name {
			return true
		}
	}
	return false
}

func (p *plugin) UpdateUpstream(original, desired *v1.Upstream) (bool, error) {
	return UpdateUpstream(original, desired)
}

func UpdateUpstream(original, desired *v1.Upstream) (didChange bool, err error) {
	originalSpec, ok := original.GetUpstreamType().(*v1.Upstream_Kube)
	if !ok {
		return false, errors.Errorf("internal error: expected *v1.Upstream_Kube, got %v", reflect.TypeOf(original.GetUpstreamType()).Name())
	}
	desiredSpec, ok := desired.GetUpstreamType().(*v1.Upstream_Kube)
	if !ok {
		return false, errors.Errorf("internal error: expected *v1.Upstream_Kube, got %v", reflect.TypeOf(original.GetUpstreamType()).Name())
	}
	// copy service spec, we don't want to overwrite that
	desiredSpec.Kube.ServiceSpec = originalSpec.Kube.GetServiceSpec()
	// copy labels; user may have written them over. cannot be auto-discovered
	desiredSpec.Kube.Selector = originalSpec.Kube.GetSelector()

	utils.UpdateUpstream(original, desired)

	return !upstreamsEqual(original, desired), nil
}

// we want to know if the upstreams are equal apart from their Status and Metadata
func upstreamsEqual(original, desired *v1.Upstream) bool {
	copyOriginal := *original
	copyDesired := *desired

	copyOriginal.Metadata = copyDesired.GetMetadata()
	copyOriginal.SetNamespacedStatuses(copyDesired.GetNamespacedStatuses())

	return copyOriginal.Equal(copyDesired)
}
