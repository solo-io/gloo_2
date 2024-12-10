package proxy_syncer_test

import (
	"testing"

	"github.com/golang/protobuf/ptypes/duration"
	"github.com/golang/protobuf/ptypes/wrappers"
	. "github.com/solo-io/gloo/projects/gateway2/proxy_syncer"
	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	networkingv1alpha3 "istio.io/api/networking/v1alpha3"
	networkingclient "istio.io/client-go/pkg/apis/networking/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestApplyDestRulesForUpstream(t *testing.T) {
	destRule := &networkingclient.DestinationRule{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "networking.istio.io/v1alpha3",
			Kind:       "DestinationRule",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: "do-failover",
		},
		Spec: networkingv1alpha3.DestinationRule{
			Host: "reviews.gwtest.svc.cluster.local",
			TrafficPolicy: &networkingv1alpha3.TrafficPolicy{
				OutlierDetection: &networkingv1alpha3.OutlierDetection{
					Consecutive_5XxErrors: &wrappers.UInt32Value{Value: 7},
					Interval:              &duration.Duration{Seconds: 300}, // 5 minutes
					BaseEjectionTime:      &duration.Duration{Seconds: 900}, // 15 minutes
				},
				LoadBalancer: &networkingv1alpha3.LoadBalancerSettings{
					LocalityLbSetting: &networkingv1alpha3.LocalityLoadBalancerSetting{
						FailoverPriority: []string{
							"topology.kubernetes.io/region",
						},
					},
				},
			},
		},
	}
	u := &gloov1.Upstream{}
	u, name := ApplyDestRulesForUpstream(&DestinationRuleWrapper{destRule}, u)
	if name == "" {
		t.Errorf("expected name to be set")
	}
	if u.OutlierDetection == nil {
		t.Fatal("expected outlier detection to be set")
	}
	if u.OutlierDetection.Consecutive_5Xx.GetValue() != 7 {
		t.Errorf("expected consecutive 5xx errors to be set")
	}
	if u.OutlierDetection.Interval.Seconds != 300 {
		t.Errorf("expected interval to be set")
	}
	if u.OutlierDetection.BaseEjectionTime.Seconds != 900 {
		t.Errorf("expected base ejection time to be set")
	}
}
