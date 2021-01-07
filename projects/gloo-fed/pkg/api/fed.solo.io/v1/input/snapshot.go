// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./snapshot.go -destination mocks/snapshot.go

// The Input Snapshot contains the set of all:
// * Services
// * Pods
// * Deployments
// * DaemonSets
// * Gateways
// * VirtualServices
// * RouteTables
// * Upstreams
// * UpstreamGroups
// * Settings
// * Proxies
// * AuthConfigs
// read from a given cluster or set of clusters, across all namespaces.
//
// A snapshot can be constructed from either a single Manager (for a single cluster)
// or a ClusterWatcher (for multiple clusters) using the SnapshotBuilder.
//
// Resources in a MultiCluster snapshot will have their ClusterName set to the
// name of the cluster from which the resource was read.

package input

import (
	"context"
	"encoding/json"

	"github.com/solo-io/skv2/pkg/verifier"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/hashicorp/go-multierror"

	"github.com/solo-io/skv2/pkg/controllerutils"
	"github.com/solo-io/skv2/pkg/multicluster"
	"sigs.k8s.io/controller-runtime/pkg/client"

	v1 "github.com/solo-io/external-apis/pkg/api/k8s/core/v1"
	v1_sets "github.com/solo-io/external-apis/pkg/api/k8s/core/v1/sets"

	apps_v1 "github.com/solo-io/external-apis/pkg/api/k8s/apps/v1"
	apps_v1_sets "github.com/solo-io/external-apis/pkg/api/k8s/apps/v1/sets"

	gateway_solo_io_v1 "github.com/solo-io/solo-apis/pkg/api/gateway.solo.io/v1"
	gateway_solo_io_v1_sets "github.com/solo-io/solo-apis/pkg/api/gateway.solo.io/v1/sets"

	gloo_solo_io_v1 "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1"
	gloo_solo_io_v1_sets "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/sets"

	enterprise_gloo_solo_io_v1 "github.com/solo-io/solo-apis/pkg/api/enterprise.gloo.solo.io/v1"
	enterprise_gloo_solo_io_v1_sets "github.com/solo-io/solo-apis/pkg/api/enterprise.gloo.solo.io/v1/sets"
)

// the snapshot of input resources consumed by translation
type Snapshot interface {

	// return the set of input Services
	Services() v1_sets.ServiceSet
	// return the set of input Pods
	Pods() v1_sets.PodSet

	// return the set of input Deployments
	Deployments() apps_v1_sets.DeploymentSet
	// return the set of input DaemonSets
	DaemonSets() apps_v1_sets.DaemonSetSet

	// return the set of input Gateways
	Gateways() gateway_solo_io_v1_sets.GatewaySet
	// return the set of input VirtualServices
	VirtualServices() gateway_solo_io_v1_sets.VirtualServiceSet
	// return the set of input RouteTables
	RouteTables() gateway_solo_io_v1_sets.RouteTableSet

	// return the set of input Upstreams
	Upstreams() gloo_solo_io_v1_sets.UpstreamSet
	// return the set of input UpstreamGroups
	UpstreamGroups() gloo_solo_io_v1_sets.UpstreamGroupSet
	// return the set of input Settings
	Settings() gloo_solo_io_v1_sets.SettingsSet
	// return the set of input Proxies
	Proxies() gloo_solo_io_v1_sets.ProxySet

	// return the set of input AuthConfigs
	AuthConfigs() enterprise_gloo_solo_io_v1_sets.AuthConfigSet
	// update the status of all input objects which support
	// the Status subresource (across multiple clusters)
	SyncStatusesMultiCluster(ctx context.Context, mcClient multicluster.Client, opts SyncStatusOptions) error
	// serialize the entire snapshot as JSON
	MarshalJSON() ([]byte, error)
}

// options for syncing input object statuses
type SyncStatusOptions struct {

	// sync status of Service objects
	Service bool
	// sync status of Pod objects
	Pod bool

	// sync status of Deployment objects
	Deployment bool
	// sync status of DaemonSet objects
	DaemonSet bool

	// sync status of Gateway objects
	Gateway bool
	// sync status of VirtualService objects
	VirtualService bool
	// sync status of RouteTable objects
	RouteTable bool

	// sync status of Upstream objects
	Upstream bool
	// sync status of UpstreamGroup objects
	UpstreamGroup bool
	// sync status of Settings objects
	Settings bool
	// sync status of Proxy objects
	Proxy bool

	// sync status of AuthConfig objects
	AuthConfig bool
}

type snapshot struct {
	name string

	services v1_sets.ServiceSet
	pods     v1_sets.PodSet

	deployments apps_v1_sets.DeploymentSet
	daemonSets  apps_v1_sets.DaemonSetSet

	gateways        gateway_solo_io_v1_sets.GatewaySet
	virtualServices gateway_solo_io_v1_sets.VirtualServiceSet
	routeTables     gateway_solo_io_v1_sets.RouteTableSet

	upstreams      gloo_solo_io_v1_sets.UpstreamSet
	upstreamGroups gloo_solo_io_v1_sets.UpstreamGroupSet
	settings       gloo_solo_io_v1_sets.SettingsSet
	proxies        gloo_solo_io_v1_sets.ProxySet

	authConfigs enterprise_gloo_solo_io_v1_sets.AuthConfigSet
}

func NewSnapshot(
	name string,

	services v1_sets.ServiceSet,
	pods v1_sets.PodSet,

	deployments apps_v1_sets.DeploymentSet,
	daemonSets apps_v1_sets.DaemonSetSet,

	gateways gateway_solo_io_v1_sets.GatewaySet,
	virtualServices gateway_solo_io_v1_sets.VirtualServiceSet,
	routeTables gateway_solo_io_v1_sets.RouteTableSet,

	upstreams gloo_solo_io_v1_sets.UpstreamSet,
	upstreamGroups gloo_solo_io_v1_sets.UpstreamGroupSet,
	settings gloo_solo_io_v1_sets.SettingsSet,
	proxies gloo_solo_io_v1_sets.ProxySet,

	authConfigs enterprise_gloo_solo_io_v1_sets.AuthConfigSet,

) Snapshot {
	return &snapshot{
		name: name,

		services:        services,
		pods:            pods,
		deployments:     deployments,
		daemonSets:      daemonSets,
		gateways:        gateways,
		virtualServices: virtualServices,
		routeTables:     routeTables,
		upstreams:       upstreams,
		upstreamGroups:  upstreamGroups,
		settings:        settings,
		proxies:         proxies,
		authConfigs:     authConfigs,
	}
}

func (s snapshot) Services() v1_sets.ServiceSet {
	return s.services
}

func (s snapshot) Pods() v1_sets.PodSet {
	return s.pods
}

func (s snapshot) Deployments() apps_v1_sets.DeploymentSet {
	return s.deployments
}

func (s snapshot) DaemonSets() apps_v1_sets.DaemonSetSet {
	return s.daemonSets
}

func (s snapshot) Gateways() gateway_solo_io_v1_sets.GatewaySet {
	return s.gateways
}

func (s snapshot) VirtualServices() gateway_solo_io_v1_sets.VirtualServiceSet {
	return s.virtualServices
}

func (s snapshot) RouteTables() gateway_solo_io_v1_sets.RouteTableSet {
	return s.routeTables
}

func (s snapshot) Upstreams() gloo_solo_io_v1_sets.UpstreamSet {
	return s.upstreams
}

func (s snapshot) UpstreamGroups() gloo_solo_io_v1_sets.UpstreamGroupSet {
	return s.upstreamGroups
}

func (s snapshot) Settings() gloo_solo_io_v1_sets.SettingsSet {
	return s.settings
}

func (s snapshot) Proxies() gloo_solo_io_v1_sets.ProxySet {
	return s.proxies
}

func (s snapshot) AuthConfigs() enterprise_gloo_solo_io_v1_sets.AuthConfigSet {
	return s.authConfigs
}

func (s snapshot) SyncStatusesMultiCluster(ctx context.Context, mcClient multicluster.Client, opts SyncStatusOptions) error {
	var errs error

	if opts.Gateway {
		for _, obj := range s.Gateways().List() {
			clusterClient, err := mcClient.Cluster(obj.ClusterName)
			if err != nil {
				errs = multierror.Append(errs, err)
				continue
			}
			if _, err := controllerutils.UpdateStatus(ctx, clusterClient, obj); err != nil {
				errs = multierror.Append(errs, err)
			}
		}
	}
	if opts.VirtualService {
		for _, obj := range s.VirtualServices().List() {
			clusterClient, err := mcClient.Cluster(obj.ClusterName)
			if err != nil {
				errs = multierror.Append(errs, err)
				continue
			}
			if _, err := controllerutils.UpdateStatus(ctx, clusterClient, obj); err != nil {
				errs = multierror.Append(errs, err)
			}
		}
	}
	if opts.RouteTable {
		for _, obj := range s.RouteTables().List() {
			clusterClient, err := mcClient.Cluster(obj.ClusterName)
			if err != nil {
				errs = multierror.Append(errs, err)
				continue
			}
			if _, err := controllerutils.UpdateStatus(ctx, clusterClient, obj); err != nil {
				errs = multierror.Append(errs, err)
			}
		}
	}

	if opts.Upstream {
		for _, obj := range s.Upstreams().List() {
			clusterClient, err := mcClient.Cluster(obj.ClusterName)
			if err != nil {
				errs = multierror.Append(errs, err)
				continue
			}
			if _, err := controllerutils.UpdateStatus(ctx, clusterClient, obj); err != nil {
				errs = multierror.Append(errs, err)
			}
		}
	}
	if opts.UpstreamGroup {
		for _, obj := range s.UpstreamGroups().List() {
			clusterClient, err := mcClient.Cluster(obj.ClusterName)
			if err != nil {
				errs = multierror.Append(errs, err)
				continue
			}
			if _, err := controllerutils.UpdateStatus(ctx, clusterClient, obj); err != nil {
				errs = multierror.Append(errs, err)
			}
		}
	}
	if opts.Settings {
		for _, obj := range s.Settings().List() {
			clusterClient, err := mcClient.Cluster(obj.ClusterName)
			if err != nil {
				errs = multierror.Append(errs, err)
				continue
			}
			if _, err := controllerutils.UpdateStatus(ctx, clusterClient, obj); err != nil {
				errs = multierror.Append(errs, err)
			}
		}
	}
	if opts.Proxy {
		for _, obj := range s.Proxies().List() {
			clusterClient, err := mcClient.Cluster(obj.ClusterName)
			if err != nil {
				errs = multierror.Append(errs, err)
				continue
			}
			if _, err := controllerutils.UpdateStatus(ctx, clusterClient, obj); err != nil {
				errs = multierror.Append(errs, err)
			}
		}
	}

	if opts.AuthConfig {
		for _, obj := range s.AuthConfigs().List() {
			clusterClient, err := mcClient.Cluster(obj.ClusterName)
			if err != nil {
				errs = multierror.Append(errs, err)
				continue
			}
			if _, err := controllerutils.UpdateStatus(ctx, clusterClient, obj); err != nil {
				errs = multierror.Append(errs, err)
			}
		}
	}
	return errs
}

func (s snapshot) MarshalJSON() ([]byte, error) {
	snapshotMap := map[string]interface{}{"name": s.name}

	snapshotMap["services"] = s.services.List()
	snapshotMap["pods"] = s.pods.List()
	snapshotMap["deployments"] = s.deployments.List()
	snapshotMap["daemonSets"] = s.daemonSets.List()
	snapshotMap["gateways"] = s.gateways.List()
	snapshotMap["virtualServices"] = s.virtualServices.List()
	snapshotMap["routeTables"] = s.routeTables.List()
	snapshotMap["upstreams"] = s.upstreams.List()
	snapshotMap["upstreamGroups"] = s.upstreamGroups.List()
	snapshotMap["settings"] = s.settings.List()
	snapshotMap["proxies"] = s.proxies.List()
	snapshotMap["authConfigs"] = s.authConfigs.List()
	return json.Marshal(snapshotMap)
}

// builds the input snapshot from API Clients.
type Builder interface {
	BuildSnapshot(ctx context.Context, name string, opts BuildOptions) (Snapshot, error)
}

// Options for building a snapshot
type BuildOptions struct {

	// List options for composing a snapshot from Services
	Services ResourceBuildOptions
	// List options for composing a snapshot from Pods
	Pods ResourceBuildOptions

	// List options for composing a snapshot from Deployments
	Deployments ResourceBuildOptions
	// List options for composing a snapshot from DaemonSets
	DaemonSets ResourceBuildOptions

	// List options for composing a snapshot from Gateways
	Gateways ResourceBuildOptions
	// List options for composing a snapshot from VirtualServices
	VirtualServices ResourceBuildOptions
	// List options for composing a snapshot from RouteTables
	RouteTables ResourceBuildOptions

	// List options for composing a snapshot from Upstreams
	Upstreams ResourceBuildOptions
	// List options for composing a snapshot from UpstreamGroups
	UpstreamGroups ResourceBuildOptions
	// List options for composing a snapshot from Settings
	Settings ResourceBuildOptions
	// List options for composing a snapshot from Proxies
	Proxies ResourceBuildOptions

	// List options for composing a snapshot from AuthConfigs
	AuthConfigs ResourceBuildOptions
}

// Options for reading resources of a given type
type ResourceBuildOptions struct {

	// List options for composing a snapshot from a resource type
	ListOptions []client.ListOption

	// If provided, ensure the resource has been verified before adding it to snapshots
	Verifier verifier.ServerResourceVerifier
}

// build a snapshot from resources across multiple clusters
type multiClusterBuilder struct {
	clusters multicluster.Interface
	client   multicluster.Client
}

// Produces snapshots of resources across all clusters defined in the ClusterSet
func NewMultiClusterBuilder(
	clusters multicluster.Interface,
	client multicluster.Client,
) Builder {
	return &multiClusterBuilder{
		clusters: clusters,
		client:   client,
	}
}

func (b *multiClusterBuilder) BuildSnapshot(ctx context.Context, name string, opts BuildOptions) (Snapshot, error) {

	services := v1_sets.NewServiceSet()
	pods := v1_sets.NewPodSet()

	deployments := apps_v1_sets.NewDeploymentSet()
	daemonSets := apps_v1_sets.NewDaemonSetSet()

	gateways := gateway_solo_io_v1_sets.NewGatewaySet()
	virtualServices := gateway_solo_io_v1_sets.NewVirtualServiceSet()
	routeTables := gateway_solo_io_v1_sets.NewRouteTableSet()

	upstreams := gloo_solo_io_v1_sets.NewUpstreamSet()
	upstreamGroups := gloo_solo_io_v1_sets.NewUpstreamGroupSet()
	settings := gloo_solo_io_v1_sets.NewSettingsSet()
	proxies := gloo_solo_io_v1_sets.NewProxySet()

	authConfigs := enterprise_gloo_solo_io_v1_sets.NewAuthConfigSet()

	var errs error

	for _, cluster := range b.clusters.ListClusters() {

		if err := b.insertServicesFromCluster(ctx, cluster, services, opts.Services); err != nil {
			errs = multierror.Append(errs, err)
		}
		if err := b.insertPodsFromCluster(ctx, cluster, pods, opts.Pods); err != nil {
			errs = multierror.Append(errs, err)
		}
		if err := b.insertDeploymentsFromCluster(ctx, cluster, deployments, opts.Deployments); err != nil {
			errs = multierror.Append(errs, err)
		}
		if err := b.insertDaemonSetsFromCluster(ctx, cluster, daemonSets, opts.DaemonSets); err != nil {
			errs = multierror.Append(errs, err)
		}
		if err := b.insertGatewaysFromCluster(ctx, cluster, gateways, opts.Gateways); err != nil {
			errs = multierror.Append(errs, err)
		}
		if err := b.insertVirtualServicesFromCluster(ctx, cluster, virtualServices, opts.VirtualServices); err != nil {
			errs = multierror.Append(errs, err)
		}
		if err := b.insertRouteTablesFromCluster(ctx, cluster, routeTables, opts.RouteTables); err != nil {
			errs = multierror.Append(errs, err)
		}
		if err := b.insertUpstreamsFromCluster(ctx, cluster, upstreams, opts.Upstreams); err != nil {
			errs = multierror.Append(errs, err)
		}
		if err := b.insertUpstreamGroupsFromCluster(ctx, cluster, upstreamGroups, opts.UpstreamGroups); err != nil {
			errs = multierror.Append(errs, err)
		}
		if err := b.insertSettingsFromCluster(ctx, cluster, settings, opts.Settings); err != nil {
			errs = multierror.Append(errs, err)
		}
		if err := b.insertProxiesFromCluster(ctx, cluster, proxies, opts.Proxies); err != nil {
			errs = multierror.Append(errs, err)
		}
		if err := b.insertAuthConfigsFromCluster(ctx, cluster, authConfigs, opts.AuthConfigs); err != nil {
			errs = multierror.Append(errs, err)
		}

	}

	outputSnap := NewSnapshot(
		name,

		services,
		pods,
		deployments,
		daemonSets,
		gateways,
		virtualServices,
		routeTables,
		upstreams,
		upstreamGroups,
		settings,
		proxies,
		authConfigs,
	)

	return outputSnap, errs
}

func (b *multiClusterBuilder) insertServicesFromCluster(ctx context.Context, cluster string, services v1_sets.ServiceSet, opts ResourceBuildOptions) error {
	serviceClient, err := v1.NewMulticlusterServiceClient(b.client).Cluster(cluster)
	if err != nil {
		return err
	}

	if opts.Verifier != nil {
		mgr, err := b.clusters.Cluster(cluster)
		if err != nil {
			return err
		}

		gvk := schema.GroupVersionKind{
			Group:   "",
			Version: "v1",
			Kind:    "Service",
		}

		if resourceRegistered, err := opts.Verifier.VerifyServerResource(
			cluster,
			mgr.GetConfig(),
			gvk,
		); err != nil {
			return err
		} else if !resourceRegistered {
			return nil
		}
	}

	serviceList, err := serviceClient.ListService(ctx, opts.ListOptions...)
	if err != nil {
		return err
	}

	for _, item := range serviceList.Items {
		item := item               // pike
		item.ClusterName = cluster // set cluster for in-memory processing
		services.Insert(&item)
	}

	return nil
}
func (b *multiClusterBuilder) insertPodsFromCluster(ctx context.Context, cluster string, pods v1_sets.PodSet, opts ResourceBuildOptions) error {
	podClient, err := v1.NewMulticlusterPodClient(b.client).Cluster(cluster)
	if err != nil {
		return err
	}

	if opts.Verifier != nil {
		mgr, err := b.clusters.Cluster(cluster)
		if err != nil {
			return err
		}

		gvk := schema.GroupVersionKind{
			Group:   "",
			Version: "v1",
			Kind:    "Pod",
		}

		if resourceRegistered, err := opts.Verifier.VerifyServerResource(
			cluster,
			mgr.GetConfig(),
			gvk,
		); err != nil {
			return err
		} else if !resourceRegistered {
			return nil
		}
	}

	podList, err := podClient.ListPod(ctx, opts.ListOptions...)
	if err != nil {
		return err
	}

	for _, item := range podList.Items {
		item := item               // pike
		item.ClusterName = cluster // set cluster for in-memory processing
		pods.Insert(&item)
	}

	return nil
}

func (b *multiClusterBuilder) insertDeploymentsFromCluster(ctx context.Context, cluster string, deployments apps_v1_sets.DeploymentSet, opts ResourceBuildOptions) error {
	deploymentClient, err := apps_v1.NewMulticlusterDeploymentClient(b.client).Cluster(cluster)
	if err != nil {
		return err
	}

	if opts.Verifier != nil {
		mgr, err := b.clusters.Cluster(cluster)
		if err != nil {
			return err
		}

		gvk := schema.GroupVersionKind{
			Group:   "apps",
			Version: "v1",
			Kind:    "Deployment",
		}

		if resourceRegistered, err := opts.Verifier.VerifyServerResource(
			cluster,
			mgr.GetConfig(),
			gvk,
		); err != nil {
			return err
		} else if !resourceRegistered {
			return nil
		}
	}

	deploymentList, err := deploymentClient.ListDeployment(ctx, opts.ListOptions...)
	if err != nil {
		return err
	}

	for _, item := range deploymentList.Items {
		item := item               // pike
		item.ClusterName = cluster // set cluster for in-memory processing
		deployments.Insert(&item)
	}

	return nil
}
func (b *multiClusterBuilder) insertDaemonSetsFromCluster(ctx context.Context, cluster string, daemonSets apps_v1_sets.DaemonSetSet, opts ResourceBuildOptions) error {
	daemonSetClient, err := apps_v1.NewMulticlusterDaemonSetClient(b.client).Cluster(cluster)
	if err != nil {
		return err
	}

	if opts.Verifier != nil {
		mgr, err := b.clusters.Cluster(cluster)
		if err != nil {
			return err
		}

		gvk := schema.GroupVersionKind{
			Group:   "apps",
			Version: "v1",
			Kind:    "DaemonSet",
		}

		if resourceRegistered, err := opts.Verifier.VerifyServerResource(
			cluster,
			mgr.GetConfig(),
			gvk,
		); err != nil {
			return err
		} else if !resourceRegistered {
			return nil
		}
	}

	daemonSetList, err := daemonSetClient.ListDaemonSet(ctx, opts.ListOptions...)
	if err != nil {
		return err
	}

	for _, item := range daemonSetList.Items {
		item := item               // pike
		item.ClusterName = cluster // set cluster for in-memory processing
		daemonSets.Insert(&item)
	}

	return nil
}

func (b *multiClusterBuilder) insertGatewaysFromCluster(ctx context.Context, cluster string, gateways gateway_solo_io_v1_sets.GatewaySet, opts ResourceBuildOptions) error {
	gatewayClient, err := gateway_solo_io_v1.NewMulticlusterGatewayClient(b.client).Cluster(cluster)
	if err != nil {
		return err
	}

	if opts.Verifier != nil {
		mgr, err := b.clusters.Cluster(cluster)
		if err != nil {
			return err
		}

		gvk := schema.GroupVersionKind{
			Group:   "gateway.solo.io",
			Version: "v1",
			Kind:    "Gateway",
		}

		if resourceRegistered, err := opts.Verifier.VerifyServerResource(
			cluster,
			mgr.GetConfig(),
			gvk,
		); err != nil {
			return err
		} else if !resourceRegistered {
			return nil
		}
	}

	gatewayList, err := gatewayClient.ListGateway(ctx, opts.ListOptions...)
	if err != nil {
		return err
	}

	for _, item := range gatewayList.Items {
		item := item               // pike
		item.ClusterName = cluster // set cluster for in-memory processing
		gateways.Insert(&item)
	}

	return nil
}
func (b *multiClusterBuilder) insertVirtualServicesFromCluster(ctx context.Context, cluster string, virtualServices gateway_solo_io_v1_sets.VirtualServiceSet, opts ResourceBuildOptions) error {
	virtualServiceClient, err := gateway_solo_io_v1.NewMulticlusterVirtualServiceClient(b.client).Cluster(cluster)
	if err != nil {
		return err
	}

	if opts.Verifier != nil {
		mgr, err := b.clusters.Cluster(cluster)
		if err != nil {
			return err
		}

		gvk := schema.GroupVersionKind{
			Group:   "gateway.solo.io",
			Version: "v1",
			Kind:    "VirtualService",
		}

		if resourceRegistered, err := opts.Verifier.VerifyServerResource(
			cluster,
			mgr.GetConfig(),
			gvk,
		); err != nil {
			return err
		} else if !resourceRegistered {
			return nil
		}
	}

	virtualServiceList, err := virtualServiceClient.ListVirtualService(ctx, opts.ListOptions...)
	if err != nil {
		return err
	}

	for _, item := range virtualServiceList.Items {
		item := item               // pike
		item.ClusterName = cluster // set cluster for in-memory processing
		virtualServices.Insert(&item)
	}

	return nil
}
func (b *multiClusterBuilder) insertRouteTablesFromCluster(ctx context.Context, cluster string, routeTables gateway_solo_io_v1_sets.RouteTableSet, opts ResourceBuildOptions) error {
	routeTableClient, err := gateway_solo_io_v1.NewMulticlusterRouteTableClient(b.client).Cluster(cluster)
	if err != nil {
		return err
	}

	if opts.Verifier != nil {
		mgr, err := b.clusters.Cluster(cluster)
		if err != nil {
			return err
		}

		gvk := schema.GroupVersionKind{
			Group:   "gateway.solo.io",
			Version: "v1",
			Kind:    "RouteTable",
		}

		if resourceRegistered, err := opts.Verifier.VerifyServerResource(
			cluster,
			mgr.GetConfig(),
			gvk,
		); err != nil {
			return err
		} else if !resourceRegistered {
			return nil
		}
	}

	routeTableList, err := routeTableClient.ListRouteTable(ctx, opts.ListOptions...)
	if err != nil {
		return err
	}

	for _, item := range routeTableList.Items {
		item := item               // pike
		item.ClusterName = cluster // set cluster for in-memory processing
		routeTables.Insert(&item)
	}

	return nil
}

func (b *multiClusterBuilder) insertUpstreamsFromCluster(ctx context.Context, cluster string, upstreams gloo_solo_io_v1_sets.UpstreamSet, opts ResourceBuildOptions) error {
	upstreamClient, err := gloo_solo_io_v1.NewMulticlusterUpstreamClient(b.client).Cluster(cluster)
	if err != nil {
		return err
	}

	if opts.Verifier != nil {
		mgr, err := b.clusters.Cluster(cluster)
		if err != nil {
			return err
		}

		gvk := schema.GroupVersionKind{
			Group:   "gloo.solo.io",
			Version: "v1",
			Kind:    "Upstream",
		}

		if resourceRegistered, err := opts.Verifier.VerifyServerResource(
			cluster,
			mgr.GetConfig(),
			gvk,
		); err != nil {
			return err
		} else if !resourceRegistered {
			return nil
		}
	}

	upstreamList, err := upstreamClient.ListUpstream(ctx, opts.ListOptions...)
	if err != nil {
		return err
	}

	for _, item := range upstreamList.Items {
		item := item               // pike
		item.ClusterName = cluster // set cluster for in-memory processing
		upstreams.Insert(&item)
	}

	return nil
}
func (b *multiClusterBuilder) insertUpstreamGroupsFromCluster(ctx context.Context, cluster string, upstreamGroups gloo_solo_io_v1_sets.UpstreamGroupSet, opts ResourceBuildOptions) error {
	upstreamGroupClient, err := gloo_solo_io_v1.NewMulticlusterUpstreamGroupClient(b.client).Cluster(cluster)
	if err != nil {
		return err
	}

	if opts.Verifier != nil {
		mgr, err := b.clusters.Cluster(cluster)
		if err != nil {
			return err
		}

		gvk := schema.GroupVersionKind{
			Group:   "gloo.solo.io",
			Version: "v1",
			Kind:    "UpstreamGroup",
		}

		if resourceRegistered, err := opts.Verifier.VerifyServerResource(
			cluster,
			mgr.GetConfig(),
			gvk,
		); err != nil {
			return err
		} else if !resourceRegistered {
			return nil
		}
	}

	upstreamGroupList, err := upstreamGroupClient.ListUpstreamGroup(ctx, opts.ListOptions...)
	if err != nil {
		return err
	}

	for _, item := range upstreamGroupList.Items {
		item := item               // pike
		item.ClusterName = cluster // set cluster for in-memory processing
		upstreamGroups.Insert(&item)
	}

	return nil
}
func (b *multiClusterBuilder) insertSettingsFromCluster(ctx context.Context, cluster string, settings gloo_solo_io_v1_sets.SettingsSet, opts ResourceBuildOptions) error {
	settingsClient, err := gloo_solo_io_v1.NewMulticlusterSettingsClient(b.client).Cluster(cluster)
	if err != nil {
		return err
	}

	if opts.Verifier != nil {
		mgr, err := b.clusters.Cluster(cluster)
		if err != nil {
			return err
		}

		gvk := schema.GroupVersionKind{
			Group:   "gloo.solo.io",
			Version: "v1",
			Kind:    "Settings",
		}

		if resourceRegistered, err := opts.Verifier.VerifyServerResource(
			cluster,
			mgr.GetConfig(),
			gvk,
		); err != nil {
			return err
		} else if !resourceRegistered {
			return nil
		}
	}

	settingsList, err := settingsClient.ListSettings(ctx, opts.ListOptions...)
	if err != nil {
		return err
	}

	for _, item := range settingsList.Items {
		item := item               // pike
		item.ClusterName = cluster // set cluster for in-memory processing
		settings.Insert(&item)
	}

	return nil
}
func (b *multiClusterBuilder) insertProxiesFromCluster(ctx context.Context, cluster string, proxies gloo_solo_io_v1_sets.ProxySet, opts ResourceBuildOptions) error {
	proxyClient, err := gloo_solo_io_v1.NewMulticlusterProxyClient(b.client).Cluster(cluster)
	if err != nil {
		return err
	}

	if opts.Verifier != nil {
		mgr, err := b.clusters.Cluster(cluster)
		if err != nil {
			return err
		}

		gvk := schema.GroupVersionKind{
			Group:   "gloo.solo.io",
			Version: "v1",
			Kind:    "Proxy",
		}

		if resourceRegistered, err := opts.Verifier.VerifyServerResource(
			cluster,
			mgr.GetConfig(),
			gvk,
		); err != nil {
			return err
		} else if !resourceRegistered {
			return nil
		}
	}

	proxyList, err := proxyClient.ListProxy(ctx, opts.ListOptions...)
	if err != nil {
		return err
	}

	for _, item := range proxyList.Items {
		item := item               // pike
		item.ClusterName = cluster // set cluster for in-memory processing
		proxies.Insert(&item)
	}

	return nil
}

func (b *multiClusterBuilder) insertAuthConfigsFromCluster(ctx context.Context, cluster string, authConfigs enterprise_gloo_solo_io_v1_sets.AuthConfigSet, opts ResourceBuildOptions) error {
	authConfigClient, err := enterprise_gloo_solo_io_v1.NewMulticlusterAuthConfigClient(b.client).Cluster(cluster)
	if err != nil {
		return err
	}

	if opts.Verifier != nil {
		mgr, err := b.clusters.Cluster(cluster)
		if err != nil {
			return err
		}

		gvk := schema.GroupVersionKind{
			Group:   "enterprise.gloo.solo.io",
			Version: "v1",
			Kind:    "AuthConfig",
		}

		if resourceRegistered, err := opts.Verifier.VerifyServerResource(
			cluster,
			mgr.GetConfig(),
			gvk,
		); err != nil {
			return err
		} else if !resourceRegistered {
			return nil
		}
	}

	authConfigList, err := authConfigClient.ListAuthConfig(ctx, opts.ListOptions...)
	if err != nil {
		return err
	}

	for _, item := range authConfigList.Items {
		item := item               // pike
		item.ClusterName = cluster // set cluster for in-memory processing
		authConfigs.Insert(&item)
	}

	return nil
}
