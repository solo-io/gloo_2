package xdsinspection

import (
	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins/aws/ec2"
	"github.com/solo-io/gloo/projects/gloo/pkg/translator"
	"github.com/solo-io/gloo/projects/gloo/pkg/upstreams"
)

func (xd *XdsDump) GetEc2InstancesForUpstream(upstream *gloov1.Upstream) []string {
	var out []string
	if xd == nil {
		out = append(out, "use -o wide for instance details")
		return out
	}
	clusterName := upstreams.UpstreamToClusterName(upstream)
	endpointCount := 0
	for _, clusterEndpoints := range xd.Endpoints {
		if clusterEndpoints.GetClusterName() == clusterName {
			for _, lEp := range clusterEndpoints.GetEndpoints() {
				for _, ep := range lEp.GetLbEndpoints() {
					if k, ok := ep.GetMetadata().GetFilterMetadata()[translator.SoloAnnotations]; ok {
						v, ok := k.GetFields()[ec2.InstanceIdAnnotationKey]
						if ok {
							endpointCount++
							out = append(out, v.GetStringValue())
						}
					}
				}
			}
		}
	}
	if endpointCount == 0 {
		out = append(out, "no endpoints")
	}
	return out
}
