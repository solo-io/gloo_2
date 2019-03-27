// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-kit/pkg/api/v1/reconcile"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
)

// Option to copy anything from the original to the desired before writing. Return value of false means don't update
type TransitionLicenseKeyFunc func(original, desired *LicenseKey) (bool, error)

type LicenseKeyReconciler interface {
	Reconcile(namespace string, desiredResources LicenseKeyList, transition TransitionLicenseKeyFunc, opts clients.ListOpts) error
}

func licenseKeysToResources(list LicenseKeyList) resources.ResourceList {
	var resourceList resources.ResourceList
	for _, licenseKey := range list {
		resourceList = append(resourceList, licenseKey)
	}
	return resourceList
}

func NewLicenseKeyReconciler(client LicenseKeyClient) LicenseKeyReconciler {
	return &licenseKeyReconciler{
		base: reconcile.NewReconciler(client.BaseClient()),
	}
}

type licenseKeyReconciler struct {
	base reconcile.Reconciler
}

func (r *licenseKeyReconciler) Reconcile(namespace string, desiredResources LicenseKeyList, transition TransitionLicenseKeyFunc, opts clients.ListOpts) error {
	opts = opts.WithDefaults()
	opts.Ctx = contextutils.WithLogger(opts.Ctx, "licenseKey_reconciler")
	var transitionResources reconcile.TransitionResourcesFunc
	if transition != nil {
		transitionResources = func(original, desired resources.Resource) (bool, error) {
			return transition(original.(*LicenseKey), desired.(*LicenseKey))
		}
	}
	return r.base.Reconcile(namespace, licenseKeysToResources(desiredResources), transitionResources, opts)
}
