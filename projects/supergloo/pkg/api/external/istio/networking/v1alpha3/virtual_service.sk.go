// Code generated by protoc-gen-solo-kit. DO NOT EDIT.

package v1alpha3

import (
	"sort"

	"github.com/gogo/protobuf/proto"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/kube/crd"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-kit/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// TODO: modify as needed to populate additional fields
func NewVirtualService(namespace, name string) *VirtualService {
	return &VirtualService{
		Metadata: core.Metadata{
			Name:      name,
			Namespace: namespace,
		},
	}
}

func (r *VirtualService) SetStatus(status core.Status) {
	r.Status = status
}

func (r *VirtualService) SetMetadata(meta core.Metadata) {
	r.Metadata = meta
}

type VirtualServiceList []*VirtualService
type VirtualservicesByNamespace map[string]VirtualServiceList

// namespace is optional, if left empty, names can collide if the list contains more than one with the same name
func (list VirtualServiceList) Find(namespace, name string) (*VirtualService, error) {
	for _, virtualService := range list {
		if virtualService.Metadata.Name == name {
			if namespace == "" || virtualService.Metadata.Namespace == namespace {
				return virtualService, nil
			}
		}
	}
	return nil, errors.Errorf("list did not find virtualService %v.%v", namespace, name)
}

func (list VirtualServiceList) AsResources() resources.ResourceList {
	var ress resources.ResourceList
	for _, virtualService := range list {
		ress = append(ress, virtualService)
	}
	return ress
}

func (list VirtualServiceList) AsInputResources() resources.InputResourceList {
	var ress resources.InputResourceList
	for _, virtualService := range list {
		ress = append(ress, virtualService)
	}
	return ress
}

func (list VirtualServiceList) Names() []string {
	var names []string
	for _, virtualService := range list {
		names = append(names, virtualService.Metadata.Name)
	}
	return names
}

func (list VirtualServiceList) NamespacesDotNames() []string {
	var names []string
	for _, virtualService := range list {
		names = append(names, virtualService.Metadata.Namespace+"."+virtualService.Metadata.Name)
	}
	return names
}

func (list VirtualServiceList) Sort() VirtualServiceList {
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].Metadata.Less(list[j].Metadata)
	})
	return list
}

func (list VirtualServiceList) Clone() VirtualServiceList {
	var virtualServiceList VirtualServiceList
	for _, virtualService := range list {
		virtualServiceList = append(virtualServiceList, proto.Clone(virtualService).(*VirtualService))
	}
	return virtualServiceList
}

func (list VirtualServiceList) ByNamespace() VirtualservicesByNamespace {
	byNamespace := make(VirtualservicesByNamespace)
	for _, virtualService := range list {
		byNamespace.Add(virtualService)
	}
	return byNamespace
}

func (byNamespace VirtualservicesByNamespace) Add(virtualService ...*VirtualService) {
	for _, item := range virtualService {
		byNamespace[item.Metadata.Namespace] = append(byNamespace[item.Metadata.Namespace], item)
	}
}

func (byNamespace VirtualservicesByNamespace) Clear(namespace string) {
	delete(byNamespace, namespace)
}

func (byNamespace VirtualservicesByNamespace) List() VirtualServiceList {
	var list VirtualServiceList
	for _, virtualServiceList := range byNamespace {
		list = append(list, virtualServiceList...)
	}
	return list.Sort()
}

func (byNamespace VirtualservicesByNamespace) Clone() VirtualservicesByNamespace {
	return byNamespace.List().Clone().ByNamespace()
}

var _ resources.Resource = &VirtualService{}

// Kubernetes Adapter for VirtualService

func (o *VirtualService) GetObjectKind() schema.ObjectKind {
	t := VirtualServiceCrd.TypeMeta()
	return &t
}

func (o *VirtualService) DeepCopyObject() runtime.Object {
	return resources.Clone(o).(*VirtualService)
}

var VirtualServiceCrd = crd.NewCrd("istio.networking.v1alpha3",
	"virtualservices",
	"istio.networking.v1alpha3",
	"v1alpha3",
	"VirtualService",
	"virtualservice",
	&VirtualService{})
