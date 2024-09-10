// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/enterprise/options/jwt/jwt.proto

package jwt

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/golang/protobuf/proto"
	equality "github.com/solo-io/protoc-gen-ext/pkg/equality"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = bytes.Compare
	_ = strings.Compare
	_ = equality.Equalizer(nil)
	_ = proto.Message(nil)
)

// Equal function
func (m *JwtStagedVhostExtension) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*JwtStagedVhostExtension)
	if !ok {
		that2, ok := that.(JwtStagedVhostExtension)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetBeforeExtAuth()).(equality.Equalizer); ok {
		if !h.Equal(target.GetBeforeExtAuth()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetBeforeExtAuth(), target.GetBeforeExtAuth()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetAfterExtAuth()).(equality.Equalizer); ok {
		if !h.Equal(target.GetAfterExtAuth()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetAfterExtAuth(), target.GetAfterExtAuth()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *JwtStagedRouteProvidersExtension) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*JwtStagedRouteProvidersExtension)
	if !ok {
		that2, ok := that.(JwtStagedRouteProvidersExtension)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetBeforeExtAuth()).(equality.Equalizer); ok {
		if !h.Equal(target.GetBeforeExtAuth()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetBeforeExtAuth(), target.GetBeforeExtAuth()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetAfterExtAuth()).(equality.Equalizer); ok {
		if !h.Equal(target.GetAfterExtAuth()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetAfterExtAuth(), target.GetAfterExtAuth()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *JwtStagedRouteExtension) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*JwtStagedRouteExtension)
	if !ok {
		that2, ok := that.(JwtStagedRouteExtension)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetBeforeExtAuth()).(equality.Equalizer); ok {
		if !h.Equal(target.GetBeforeExtAuth()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetBeforeExtAuth(), target.GetBeforeExtAuth()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetAfterExtAuth()).(equality.Equalizer); ok {
		if !h.Equal(target.GetAfterExtAuth()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetAfterExtAuth(), target.GetAfterExtAuth()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *VhostExtension) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*VhostExtension)
	if !ok {
		that2, ok := that.(VhostExtension)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if len(m.GetProviders()) != len(target.GetProviders()) {
		return false
	}
	for k, v := range m.GetProviders() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetProviders()[k]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetProviders()[k]) {
				return false
			}
		}

	}

	if m.GetAllowMissingOrFailedJwt() != target.GetAllowMissingOrFailedJwt() {
		return false
	}

	if m.GetValidationPolicy() != target.GetValidationPolicy() {
		return false
	}

	return true
}

// Equal function
func (m *RouteExtension) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*RouteExtension)
	if !ok {
		that2, ok := that.(RouteExtension)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if m.GetDisable() != target.GetDisable() {
		return false
	}

	return true
}

// Equal function
func (m *Provider) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Provider)
	if !ok {
		that2, ok := that.(Provider)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetJwks()).(equality.Equalizer); ok {
		if !h.Equal(target.GetJwks()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetJwks(), target.GetJwks()) {
			return false
		}
	}

	if len(m.GetAudiences()) != len(target.GetAudiences()) {
		return false
	}
	for idx, v := range m.GetAudiences() {

		if strings.Compare(v, target.GetAudiences()[idx]) != 0 {
			return false
		}

	}

	if strings.Compare(m.GetIssuer(), target.GetIssuer()) != 0 {
		return false
	}

	if h, ok := interface{}(m.GetTokenSource()).(equality.Equalizer); ok {
		if !h.Equal(target.GetTokenSource()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetTokenSource(), target.GetTokenSource()) {
			return false
		}
	}

	if m.GetKeepToken() != target.GetKeepToken() {
		return false
	}

	if len(m.GetClaimsToHeaders()) != len(target.GetClaimsToHeaders()) {
		return false
	}
	for idx, v := range m.GetClaimsToHeaders() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetClaimsToHeaders()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetClaimsToHeaders()[idx]) {
				return false
			}
		}

	}

	if h, ok := interface{}(m.GetClockSkewSeconds()).(equality.Equalizer); ok {
		if !h.Equal(target.GetClockSkewSeconds()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetClockSkewSeconds(), target.GetClockSkewSeconds()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *Jwks) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Jwks)
	if !ok {
		that2, ok := that.(Jwks)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	switch m.Jwks.(type) {

	case *Jwks_Remote:
		if _, ok := target.Jwks.(*Jwks_Remote); !ok {
			return false
		}

		if h, ok := interface{}(m.GetRemote()).(equality.Equalizer); ok {
			if !h.Equal(target.GetRemote()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetRemote(), target.GetRemote()) {
				return false
			}
		}

	case *Jwks_Local:
		if _, ok := target.Jwks.(*Jwks_Local); !ok {
			return false
		}

		if h, ok := interface{}(m.GetLocal()).(equality.Equalizer); ok {
			if !h.Equal(target.GetLocal()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetLocal(), target.GetLocal()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.Jwks != target.Jwks {
			return false
		}
	}

	return true
}

// Equal function
func (m *RemoteJwks) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*RemoteJwks)
	if !ok {
		that2, ok := that.(RemoteJwks)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetUrl(), target.GetUrl()) != 0 {
		return false
	}

	if h, ok := interface{}(m.GetUpstreamRef()).(equality.Equalizer); ok {
		if !h.Equal(target.GetUpstreamRef()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetUpstreamRef(), target.GetUpstreamRef()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetCacheDuration()).(equality.Equalizer); ok {
		if !h.Equal(target.GetCacheDuration()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetCacheDuration(), target.GetCacheDuration()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetAsyncFetch()).(equality.Equalizer); ok {
		if !h.Equal(target.GetAsyncFetch()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetAsyncFetch(), target.GetAsyncFetch()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *LocalJwks) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*LocalJwks)
	if !ok {
		that2, ok := that.(LocalJwks)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetKey(), target.GetKey()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *TokenSource) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*TokenSource)
	if !ok {
		that2, ok := that.(TokenSource)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if len(m.GetHeaders()) != len(target.GetHeaders()) {
		return false
	}
	for idx, v := range m.GetHeaders() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetHeaders()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetHeaders()[idx]) {
				return false
			}
		}

	}

	if len(m.GetQueryParams()) != len(target.GetQueryParams()) {
		return false
	}
	for idx, v := range m.GetQueryParams() {

		if strings.Compare(v, target.GetQueryParams()[idx]) != 0 {
			return false
		}

	}

	return true
}

// Equal function
func (m *ClaimToHeader) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ClaimToHeader)
	if !ok {
		that2, ok := that.(ClaimToHeader)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetClaim(), target.GetClaim()) != 0 {
		return false
	}

	if strings.Compare(m.GetHeader(), target.GetHeader()) != 0 {
		return false
	}

	if m.GetAppend() != target.GetAppend() {
		return false
	}

	return true
}

// Equal function
func (m *TokenSource_HeaderSource) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*TokenSource_HeaderSource)
	if !ok {
		that2, ok := that.(TokenSource_HeaderSource)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetHeader(), target.GetHeader()) != 0 {
		return false
	}

	if strings.Compare(m.GetPrefix(), target.GetPrefix()) != 0 {
		return false
	}

	return true
}
