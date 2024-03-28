package types

import (
	"fmt"
	"os"

	"github.com/solo-io/gloo/install/helm/gloo/generate"
	"github.com/solo-io/gloo/test/setup/defaults"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

type (
	// This follows the values schema for the Gloo Edge Helm chart
	GlooEdge generate.HelmConfig

	Image generate.Image

	App struct {
		ConfigMaps     []*corev1.ConfigMap    `yaml:"configMaps,omitempty" json:"configMaps,omitempty"`
		Service        *corev1.Service        `yaml:"service,omitempty" json:"service,omitempty"`
		ServiceAccount *corev1.ServiceAccount `yaml:"serviceAccount,omitempty" json:"serviceAccount,omitempty"`
		Deployment     *appsv1.Deployment     `yaml:"deployment,omitempty" json:"deployment,omitempty"`
		Versions       []string               `yaml:"versions,omitempty" json:"versions,omitempty"`
	}
)

func (p GlooEdge) Images() (imageRefs []string) {

	defaultImage := &generate.Image{}
	if p.Global != nil {
		if p.Global.Image != nil {
			// load override
			defaultImage = p.Global.Image
		}

		glooMtlsEnabled := p.Global.GlooMtls != nil && *p.Global.GlooMtls.Enabled
		istioSDSEnabled := p.Global.IstioSDS != nil && *p.Global.IstioSDS.Enabled

		// Check if either GlooMtls or IstioSDS is enabled.
		// Note: Image overrides are only defined on GlooMtls.
		if glooMtlsEnabled || istioSDSEnabled {
			// Add Sds image IstioProxy override if GlooMtls is enabled
			if p.Global.GlooMtls != nil && p.Global.GlooMtls.IstioProxy.Image != nil {
				imageRefs = append(imageRefs, getImageRef(p.Global.GlooMtls.Sds.Image, defaults.DefaultSdsImageName))
			} else {
				// Add default global image if GlooMtls image overrides are not set
				imageRefs = append(imageRefs, getImageRef(defaultImage, defaults.DefaultSdsImageName))
			}

			// Add EnvoySidecar image override if GlooMtls is enabled
			if p.Global.GlooMtls != nil && p.Global.GlooMtls.EnvoySidecar.Image != nil {
				imageRefs = append(imageRefs, getImageRef(p.Global.GlooMtls.EnvoySidecar.Image, defaults.DefaultGatewayImageName))
			} else {
				// Add default global image if GlooMtls image overrides are not set
				imageRefs = append(imageRefs, getImageRef(defaultImage, defaults.DefaultGatewayImageName))
			}

			// Add IstioProxy image override if IstioSDS is enabled
			if istioSDSEnabled {
				if p.Global.GlooMtls != nil && p.Global.GlooMtls.IstioProxy.Image != nil {
					imageRefs = append(imageRefs, getImageRef(p.Global.GlooMtls.IstioProxy.Image, defaults.DefaultIstioProxyImageName))
				} else {
					// Add default image
					imageRefs = append(imageRefs, fmt.Sprintf("%s/%s:%s", defaults.DefaultIstioImageRegistry, defaults.DefaultIstioProxyImageName, defaults.DefaultIstioTag))
				}
			}
		} else if istioSDSEnabled {
			// Add default global image if IstioSDS is enabled, but GlooMtls image overrides are not set
			imageRefs = append(imageRefs, getImageRef(defaultImage, defaults.DefaultSdsImageName))
			imageRefs = append(imageRefs, fmt.Sprintf("%s/%s:%s", defaults.DefaultIstioImageRegistry, defaults.DefaultIstioProxyImageName, defaults.DefaultIstioTag))
		}

	}

	// gloo image is configured via Global.Image
	imageRefs = append(imageRefs, getImageRef(defaultImage, defaults.DefaultGlooImageName))

	if p.Gateway != nil && *p.Gateway.Enabled {
		// use overrides
		if *p.Gateway.CertGenJob.Enabled {
			if p.Gateway.CertGenJob.Image != nil {
				imageRefs = append(imageRefs, getImageRef(p.Gateway.CertGenJob.Image, defaults.DefaultCertGenJobImageName))
			} else {
				imageRefs = append(imageRefs, getImageRef(defaultImage, defaults.DefaultCertGenJobImageName))
			}
		}
		if *p.Gateway.RolloutJob.Enabled {
			if p.Gateway.RolloutJob.Image != nil {
				imageRefs = append(imageRefs, getImageRef(p.Gateway.RolloutJob.Image, defaults.DefaultRolloutJobImageName))
			} else {
				imageRefs = append(imageRefs, getImageRef(defaultImage, defaults.DefaultRolloutJobImageName))
			}
		}
		if *p.Gateway.CleanupJob.Enabled {
			if p.Gateway.CleanupJob.Image != nil {
				imageRefs = append(imageRefs, getImageRef(p.Gateway.CleanupJob.Image, defaults.DefaultCleanupJobImageName))
			} else {
				imageRefs = append(imageRefs, getImageRef(defaultImage, defaults.DefaultCleanupJobImageName))
			}
		}
	} else {
		// use default global image
		imageRefs = append(imageRefs, getImageRef(defaultImage, defaults.DefaultGatewayImageName))
		imageRefs = append(imageRefs, getImageRef(defaultImage, defaults.DefaultCertGenJobImageName))
		imageRefs = append(imageRefs, getImageRef(defaultImage, defaults.DefaultRolloutJobImageName))
		imageRefs = append(imageRefs, getImageRef(defaultImage, defaults.DefaultCleanupJobImageName))
	}

	if *p.Discovery.Enabled {
		if p.Discovery.Deployment.Image != nil {
			imageRefs = append(imageRefs, getImageRef(p.Discovery.Deployment.Image, defaults.DefaultDiscoveryImageName))
		} else {
			// use default global image
			imageRefs = append(imageRefs, getImageRef(defaultImage, defaults.DefaultDiscoveryImageName))
		}
	}

	return imageRefs
}

func getImageRef(image *generate.Image, component string) string {
	return fmt.Sprintf(
		"%s/%s:%s",
		withDefault(image.Registry, os.Getenv("IMAGE_REGISTRY")),
		withDefault(image.Repository, component),
		withDefault(image.Tag, os.Getenv("VERSION")),
	)
}

func withDefault(value *string, defaultValue string) string {
	if value == nil {
		return defaultValue
	}
	return *value
}
