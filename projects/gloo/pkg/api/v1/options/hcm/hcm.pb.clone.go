// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/hcm/hcm.proto

package hcm

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_headers "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/headers"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_protocol "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/protocol"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_protocol_upgrade "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/protocol_upgrade"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_tracing "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/tracing"

	google_golang_org_protobuf_types_known_durationpb "google.golang.org/protobuf/types/known/durationpb"

	google_golang_org_protobuf_types_known_wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = bytes.Compare
	_ = strings.Compare
	_ = clone.Cloner(nil)
	_ = proto.Message(nil)
)

// Clone function
func (m *HttpConnectionManagerSettings) Clone() proto.Message {
	var target *HttpConnectionManagerSettings
	if m == nil {
		return target
	}
	target = &HttpConnectionManagerSettings{}

	if h, ok := interface{}(m.GetSkipXffAppend()).(clone.Cloner); ok {
		target.SkipXffAppend = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	} else {
		target.SkipXffAppend = proto.Clone(m.GetSkipXffAppend()).(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	}

	if h, ok := interface{}(m.GetVia()).(clone.Cloner); ok {
		target.Via = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.StringValue)
	} else {
		target.Via = proto.Clone(m.GetVia()).(*google_golang_org_protobuf_types_known_wrapperspb.StringValue)
	}

	if h, ok := interface{}(m.GetXffNumTrustedHops()).(clone.Cloner); ok {
		target.XffNumTrustedHops = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.UInt32Value)
	} else {
		target.XffNumTrustedHops = proto.Clone(m.GetXffNumTrustedHops()).(*google_golang_org_protobuf_types_known_wrapperspb.UInt32Value)
	}

	if h, ok := interface{}(m.GetUseRemoteAddress()).(clone.Cloner); ok {
		target.UseRemoteAddress = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	} else {
		target.UseRemoteAddress = proto.Clone(m.GetUseRemoteAddress()).(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	}

	if h, ok := interface{}(m.GetGenerateRequestId()).(clone.Cloner); ok {
		target.GenerateRequestId = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	} else {
		target.GenerateRequestId = proto.Clone(m.GetGenerateRequestId()).(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	}

	if h, ok := interface{}(m.GetProxy_100Continue()).(clone.Cloner); ok {
		target.Proxy_100Continue = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	} else {
		target.Proxy_100Continue = proto.Clone(m.GetProxy_100Continue()).(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	}

	if h, ok := interface{}(m.GetStreamIdleTimeout()).(clone.Cloner); ok {
		target.StreamIdleTimeout = h.Clone().(*google_golang_org_protobuf_types_known_durationpb.Duration)
	} else {
		target.StreamIdleTimeout = proto.Clone(m.GetStreamIdleTimeout()).(*google_golang_org_protobuf_types_known_durationpb.Duration)
	}

	if h, ok := interface{}(m.GetIdleTimeout()).(clone.Cloner); ok {
		target.IdleTimeout = h.Clone().(*google_golang_org_protobuf_types_known_durationpb.Duration)
	} else {
		target.IdleTimeout = proto.Clone(m.GetIdleTimeout()).(*google_golang_org_protobuf_types_known_durationpb.Duration)
	}

	if h, ok := interface{}(m.GetMaxRequestHeadersKb()).(clone.Cloner); ok {
		target.MaxRequestHeadersKb = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.UInt32Value)
	} else {
		target.MaxRequestHeadersKb = proto.Clone(m.GetMaxRequestHeadersKb()).(*google_golang_org_protobuf_types_known_wrapperspb.UInt32Value)
	}

	if h, ok := interface{}(m.GetRequestTimeout()).(clone.Cloner); ok {
		target.RequestTimeout = h.Clone().(*google_golang_org_protobuf_types_known_durationpb.Duration)
	} else {
		target.RequestTimeout = proto.Clone(m.GetRequestTimeout()).(*google_golang_org_protobuf_types_known_durationpb.Duration)
	}

	if h, ok := interface{}(m.GetRequestHeadersTimeout()).(clone.Cloner); ok {
		target.RequestHeadersTimeout = h.Clone().(*google_golang_org_protobuf_types_known_durationpb.Duration)
	} else {
		target.RequestHeadersTimeout = proto.Clone(m.GetRequestHeadersTimeout()).(*google_golang_org_protobuf_types_known_durationpb.Duration)
	}

	if h, ok := interface{}(m.GetDrainTimeout()).(clone.Cloner); ok {
		target.DrainTimeout = h.Clone().(*google_golang_org_protobuf_types_known_durationpb.Duration)
	} else {
		target.DrainTimeout = proto.Clone(m.GetDrainTimeout()).(*google_golang_org_protobuf_types_known_durationpb.Duration)
	}

	if h, ok := interface{}(m.GetDelayedCloseTimeout()).(clone.Cloner); ok {
		target.DelayedCloseTimeout = h.Clone().(*google_golang_org_protobuf_types_known_durationpb.Duration)
	} else {
		target.DelayedCloseTimeout = proto.Clone(m.GetDelayedCloseTimeout()).(*google_golang_org_protobuf_types_known_durationpb.Duration)
	}

	if h, ok := interface{}(m.GetServerName()).(clone.Cloner); ok {
		target.ServerName = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.StringValue)
	} else {
		target.ServerName = proto.Clone(m.GetServerName()).(*google_golang_org_protobuf_types_known_wrapperspb.StringValue)
	}

	if h, ok := interface{}(m.GetStripAnyHostPort()).(clone.Cloner); ok {
		target.StripAnyHostPort = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	} else {
		target.StripAnyHostPort = proto.Clone(m.GetStripAnyHostPort()).(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	}

	if h, ok := interface{}(m.GetAcceptHttp_10()).(clone.Cloner); ok {
		target.AcceptHttp_10 = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	} else {
		target.AcceptHttp_10 = proto.Clone(m.GetAcceptHttp_10()).(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	}

	if h, ok := interface{}(m.GetDefaultHostForHttp_10()).(clone.Cloner); ok {
		target.DefaultHostForHttp_10 = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.StringValue)
	} else {
		target.DefaultHostForHttp_10 = proto.Clone(m.GetDefaultHostForHttp_10()).(*google_golang_org_protobuf_types_known_wrapperspb.StringValue)
	}

	if h, ok := interface{}(m.GetAllowChunkedLength()).(clone.Cloner); ok {
		target.AllowChunkedLength = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	} else {
		target.AllowChunkedLength = proto.Clone(m.GetAllowChunkedLength()).(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	}

	if h, ok := interface{}(m.GetEnableTrailers()).(clone.Cloner); ok {
		target.EnableTrailers = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	} else {
		target.EnableTrailers = proto.Clone(m.GetEnableTrailers()).(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	}

	if h, ok := interface{}(m.GetTracing()).(clone.Cloner); ok {
		target.Tracing = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_tracing.ListenerTracingSettings)
	} else {
		target.Tracing = proto.Clone(m.GetTracing()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_tracing.ListenerTracingSettings)
	}

	target.ForwardClientCertDetails = m.GetForwardClientCertDetails()

	if h, ok := interface{}(m.GetSetCurrentClientCertDetails()).(clone.Cloner); ok {
		target.SetCurrentClientCertDetails = h.Clone().(*HttpConnectionManagerSettings_SetCurrentClientCertDetails)
	} else {
		target.SetCurrentClientCertDetails = proto.Clone(m.GetSetCurrentClientCertDetails()).(*HttpConnectionManagerSettings_SetCurrentClientCertDetails)
	}

	if h, ok := interface{}(m.GetPreserveExternalRequestId()).(clone.Cloner); ok {
		target.PreserveExternalRequestId = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	} else {
		target.PreserveExternalRequestId = proto.Clone(m.GetPreserveExternalRequestId()).(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	}

	if m.GetUpgrades() != nil {
		target.Upgrades = make([]*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_protocol_upgrade.ProtocolUpgradeConfig, len(m.GetUpgrades()))
		for idx, v := range m.GetUpgrades() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Upgrades[idx] = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_protocol_upgrade.ProtocolUpgradeConfig)
			} else {
				target.Upgrades[idx] = proto.Clone(v).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_protocol_upgrade.ProtocolUpgradeConfig)
			}

		}
	}

	if h, ok := interface{}(m.GetMaxConnectionDuration()).(clone.Cloner); ok {
		target.MaxConnectionDuration = h.Clone().(*google_golang_org_protobuf_types_known_durationpb.Duration)
	} else {
		target.MaxConnectionDuration = proto.Clone(m.GetMaxConnectionDuration()).(*google_golang_org_protobuf_types_known_durationpb.Duration)
	}

	if h, ok := interface{}(m.GetMaxStreamDuration()).(clone.Cloner); ok {
		target.MaxStreamDuration = h.Clone().(*google_golang_org_protobuf_types_known_durationpb.Duration)
	} else {
		target.MaxStreamDuration = proto.Clone(m.GetMaxStreamDuration()).(*google_golang_org_protobuf_types_known_durationpb.Duration)
	}

	if h, ok := interface{}(m.GetMaxHeadersCount()).(clone.Cloner); ok {
		target.MaxHeadersCount = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.UInt32Value)
	} else {
		target.MaxHeadersCount = proto.Clone(m.GetMaxHeadersCount()).(*google_golang_org_protobuf_types_known_wrapperspb.UInt32Value)
	}

	target.HeadersWithUnderscoresAction = m.GetHeadersWithUnderscoresAction()

	if h, ok := interface{}(m.GetMaxRequestsPerConnection()).(clone.Cloner); ok {
		target.MaxRequestsPerConnection = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.UInt32Value)
	} else {
		target.MaxRequestsPerConnection = proto.Clone(m.GetMaxRequestsPerConnection()).(*google_golang_org_protobuf_types_known_wrapperspb.UInt32Value)
	}

	target.ServerHeaderTransformation = m.GetServerHeaderTransformation()

	target.PathWithEscapedSlashesAction = m.GetPathWithEscapedSlashesAction()

	target.CodecType = m.GetCodecType()

	if h, ok := interface{}(m.GetMergeSlashes()).(clone.Cloner); ok {
		target.MergeSlashes = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	} else {
		target.MergeSlashes = proto.Clone(m.GetMergeSlashes()).(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	}

	if h, ok := interface{}(m.GetNormalizePath()).(clone.Cloner); ok {
		target.NormalizePath = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	} else {
		target.NormalizePath = proto.Clone(m.GetNormalizePath()).(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	}

	if h, ok := interface{}(m.GetUuidRequestIdConfig()).(clone.Cloner); ok {
		target.UuidRequestIdConfig = h.Clone().(*HttpConnectionManagerSettings_UuidRequestIdConfigSettings)
	} else {
		target.UuidRequestIdConfig = proto.Clone(m.GetUuidRequestIdConfig()).(*HttpConnectionManagerSettings_UuidRequestIdConfigSettings)
	}

	if h, ok := interface{}(m.GetHttp2ProtocolOptions()).(clone.Cloner); ok {
		target.Http2ProtocolOptions = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_protocol.Http2ProtocolOptions)
	} else {
		target.Http2ProtocolOptions = proto.Clone(m.GetHttp2ProtocolOptions()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_protocol.Http2ProtocolOptions)
	}

	if h, ok := interface{}(m.GetInternalAddressConfig()).(clone.Cloner); ok {
		target.InternalAddressConfig = h.Clone().(*HttpConnectionManagerSettings_InternalAddressConfig)
	} else {
		target.InternalAddressConfig = proto.Clone(m.GetInternalAddressConfig()).(*HttpConnectionManagerSettings_InternalAddressConfig)
	}

	if h, ok := interface{}(m.GetAppendXForwardedPort()).(clone.Cloner); ok {
		target.AppendXForwardedPort = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	} else {
		target.AppendXForwardedPort = proto.Clone(m.GetAppendXForwardedPort()).(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	}

	if h, ok := interface{}(m.GetEarlyHeaderManipulation()).(clone.Cloner); ok {
		target.EarlyHeaderManipulation = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_headers.HeaderManipulation)
	} else {
		target.EarlyHeaderManipulation = proto.Clone(m.GetEarlyHeaderManipulation()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_options_headers.HeaderManipulation)
	}

	switch m.HeaderFormat.(type) {

	case *HttpConnectionManagerSettings_ProperCaseHeaderKeyFormat:

		if h, ok := interface{}(m.GetProperCaseHeaderKeyFormat()).(clone.Cloner); ok {
			target.HeaderFormat = &HttpConnectionManagerSettings_ProperCaseHeaderKeyFormat{
				ProperCaseHeaderKeyFormat: h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue),
			}
		} else {
			target.HeaderFormat = &HttpConnectionManagerSettings_ProperCaseHeaderKeyFormat{
				ProperCaseHeaderKeyFormat: proto.Clone(m.GetProperCaseHeaderKeyFormat()).(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue),
			}
		}

	case *HttpConnectionManagerSettings_PreserveCaseHeaderKeyFormat:

		if h, ok := interface{}(m.GetPreserveCaseHeaderKeyFormat()).(clone.Cloner); ok {
			target.HeaderFormat = &HttpConnectionManagerSettings_PreserveCaseHeaderKeyFormat{
				PreserveCaseHeaderKeyFormat: h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue),
			}
		} else {
			target.HeaderFormat = &HttpConnectionManagerSettings_PreserveCaseHeaderKeyFormat{
				PreserveCaseHeaderKeyFormat: proto.Clone(m.GetPreserveCaseHeaderKeyFormat()).(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue),
			}
		}

	}

	return target
}

// Clone function
func (m *HttpConnectionManagerSettings_SetCurrentClientCertDetails) Clone() proto.Message {
	var target *HttpConnectionManagerSettings_SetCurrentClientCertDetails
	if m == nil {
		return target
	}
	target = &HttpConnectionManagerSettings_SetCurrentClientCertDetails{}

	if h, ok := interface{}(m.GetSubject()).(clone.Cloner); ok {
		target.Subject = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	} else {
		target.Subject = proto.Clone(m.GetSubject()).(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	}

	if h, ok := interface{}(m.GetCert()).(clone.Cloner); ok {
		target.Cert = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	} else {
		target.Cert = proto.Clone(m.GetCert()).(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	}

	if h, ok := interface{}(m.GetChain()).(clone.Cloner); ok {
		target.Chain = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	} else {
		target.Chain = proto.Clone(m.GetChain()).(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	}

	if h, ok := interface{}(m.GetDns()).(clone.Cloner); ok {
		target.Dns = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	} else {
		target.Dns = proto.Clone(m.GetDns()).(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	}

	if h, ok := interface{}(m.GetUri()).(clone.Cloner); ok {
		target.Uri = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	} else {
		target.Uri = proto.Clone(m.GetUri()).(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	}

	return target
}

// Clone function
func (m *HttpConnectionManagerSettings_UuidRequestIdConfigSettings) Clone() proto.Message {
	var target *HttpConnectionManagerSettings_UuidRequestIdConfigSettings
	if m == nil {
		return target
	}
	target = &HttpConnectionManagerSettings_UuidRequestIdConfigSettings{}

	if h, ok := interface{}(m.GetPackTraceReason()).(clone.Cloner); ok {
		target.PackTraceReason = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	} else {
		target.PackTraceReason = proto.Clone(m.GetPackTraceReason()).(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	}

	if h, ok := interface{}(m.GetUseRequestIdForTraceSampling()).(clone.Cloner); ok {
		target.UseRequestIdForTraceSampling = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	} else {
		target.UseRequestIdForTraceSampling = proto.Clone(m.GetUseRequestIdForTraceSampling()).(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	}

	return target
}

// Clone function
func (m *HttpConnectionManagerSettings_CidrRange) Clone() proto.Message {
	var target *HttpConnectionManagerSettings_CidrRange
	if m == nil {
		return target
	}
	target = &HttpConnectionManagerSettings_CidrRange{}

	target.AddressPrefix = m.GetAddressPrefix()

	if h, ok := interface{}(m.GetPrefixLen()).(clone.Cloner); ok {
		target.PrefixLen = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.UInt32Value)
	} else {
		target.PrefixLen = proto.Clone(m.GetPrefixLen()).(*google_golang_org_protobuf_types_known_wrapperspb.UInt32Value)
	}

	return target
}

// Clone function
func (m *HttpConnectionManagerSettings_InternalAddressConfig) Clone() proto.Message {
	var target *HttpConnectionManagerSettings_InternalAddressConfig
	if m == nil {
		return target
	}
	target = &HttpConnectionManagerSettings_InternalAddressConfig{}

	if h, ok := interface{}(m.GetUnixSockets()).(clone.Cloner); ok {
		target.UnixSockets = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	} else {
		target.UnixSockets = proto.Clone(m.GetUnixSockets()).(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	}

	if m.GetCidrRanges() != nil {
		target.CidrRanges = make([]*HttpConnectionManagerSettings_CidrRange, len(m.GetCidrRanges()))
		for idx, v := range m.GetCidrRanges() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.CidrRanges[idx] = h.Clone().(*HttpConnectionManagerSettings_CidrRange)
			} else {
				target.CidrRanges[idx] = proto.Clone(v).(*HttpConnectionManagerSettings_CidrRange)
			}

		}
	}

	return target
}
