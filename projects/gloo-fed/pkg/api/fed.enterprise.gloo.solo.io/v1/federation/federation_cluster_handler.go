// Code generated by skv2. DO NOT EDIT.

// Definition for federated resource cluster handler templates
package federation

import (
	"context"

	"github.com/avast/retry-go"
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/skv2/pkg/multicluster"
	fed_enterprise_gloo_solo_io_v1 "github.com/solo-io/solo-projects/projects/gloo-fed/pkg/api/fed.enterprise.gloo.solo.io/v1"
	mc_types "github.com/solo-io/solo-projects/projects/gloo-fed/pkg/api/fed.solo.io/core/v1"
	"github.com/solo-io/solo-projects/projects/gloo-fed/pkg/federation/placement"
	"go.uber.org/zap"
	"k8s.io/apimachinery/pkg/api/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

var ClusterHandlerRetryAttempts uint = 5

type clusterHandler struct {
	ctx     context.Context
	clients fed_enterprise_gloo_solo_io_v1.Clientset
	factory placement.StatusBuilderFactory
}

func NewClusterHandler(ctx context.Context, clients fed_enterprise_gloo_solo_io_v1.Clientset, factory placement.StatusBuilderFactory) multicluster.ClusterHandler {
	return &clusterHandler{
		ctx:     ctx,
		clients: clients,
		factory: factory,
	}
}

func (f *clusterHandler) AddCluster(_ context.Context, cluster string, _ manager.Manager) {
	f.handleClusterEvent(cluster)
}

func (f *clusterHandler) RemoveCluster(cluster string) {
	f.handleClusterEvent(cluster)
}

func (f *clusterHandler) handleClusterEvent(cluster string) {

	federatedAuthConfigList, err := f.clients.FederatedAuthConfigs().ListFederatedAuthConfig(f.ctx)
	if err != nil {
		contextutils.LoggerFrom(f.ctx).Errorf("Failed to list FederatedAuthConfigs referencing cluster %s", cluster)
	} else {
		for _, item := range federatedAuthConfigList.Items {
			item := item
			if err := f.maybeUpdateFederatedAuthConfigStatusWithRetries(&item, cluster); err != nil {
				contextutils.LoggerFrom(f.ctx).Errorw("Failed to update status on FederatedAuthConfig",
					zap.Error(err),
					zap.Any("FederatedAuthConfig", item))
			}
		}
	}
}

func (f *clusterHandler) maybeUpdateFederatedAuthConfigStatusWithRetries(item *fed_enterprise_gloo_solo_io_v1.FederatedAuthConfig, cluster string) error {
	return retry.Do(func() error {
		err := f.maybeUpdateFederatedAuthConfigStatus(item, cluster)
		if err != nil && errors.IsNotFound(err) {
			// If the resource no longer exists, there is nothing to do.
			return nil
		} else if err != nil && errors.IsConflict(err) {
			// On conflict, retry with the new object to pick up any changes to the resource's spec.
			obj, err := f.clients.FederatedAuthConfigs().GetFederatedAuthConfig(f.ctx, client.ObjectKey{Namespace: item.Namespace, Name: item.Name})
			if err != nil {
				return err
			}
			item = obj
		}
		return err
	}, retry.Attempts(ClusterHandlerRetryAttempts))
}

func (f *clusterHandler) maybeUpdateFederatedAuthConfigStatus(item *fed_enterprise_gloo_solo_io_v1.FederatedAuthConfig, cluster string) error {
	for _, c := range item.Spec.Placement.GetClusters() {
		if c == cluster {
			// An existing resource references the given cluster. Update its status to trigger a resync.
			item.Status.PlacementStatus = f.factory.GetBuilder().
				UpdateUnprocessed(item.Status.PlacementStatus, placement.ClusterEventTriggered(cluster), mc_types.PlacementStatus_PENDING).
				// Do not update the observed generation or written by fields as we have not actually processed the resource.
				Eject(item.Status.PlacementStatus.GetObservedGeneration())

			return f.clients.FederatedAuthConfigs().UpdateFederatedAuthConfigStatus(f.ctx, item)
		}
	}
	return nil
}
