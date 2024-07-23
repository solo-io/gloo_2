// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/kubegateway/api/v1alpha1/kube/container.proto

package kube

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
func (m *Image) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Image)
	if !ok {
		that2, ok := that.(Image)
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

	if h, ok := interface{}(m.GetRegistry()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRegistry()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRegistry(), target.GetRegistry()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetRepository()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRepository()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRepository(), target.GetRepository()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetTag()).(equality.Equalizer); ok {
		if !h.Equal(target.GetTag()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetTag(), target.GetTag()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetDigest()).(equality.Equalizer); ok {
		if !h.Equal(target.GetDigest()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetDigest(), target.GetDigest()) {
			return false
		}
	}

	if m.GetPullPolicy() != target.GetPullPolicy() {
		return false
	}

	return true
}

// Equal function
func (m *ResourceRequirements) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ResourceRequirements)
	if !ok {
		that2, ok := that.(ResourceRequirements)
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

	if len(m.GetLimits()) != len(target.GetLimits()) {
		return false
	}
	for k, v := range m.GetLimits() {

		if strings.Compare(v, target.GetLimits()[k]) != 0 {
			return false
		}

	}

	if len(m.GetRequests()) != len(target.GetRequests()) {
		return false
	}
	for k, v := range m.GetRequests() {

		if strings.Compare(v, target.GetRequests()[k]) != 0 {
			return false
		}

	}

	return true
}