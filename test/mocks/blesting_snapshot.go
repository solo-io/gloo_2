// Code generated by protoc-gen-solo-kit. DO NOT EDIT.

package mocks

import (
	"go.uber.org/zap"

	"github.com/mitchellh/hashstructure"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

type BlestingSnapshot struct {
	Fakes FakesByNamespace
}

func (s BlestingSnapshot) Clone() BlestingSnapshot {
	return BlestingSnapshot{
		Fakes: s.Fakes.Clone(),
	}
}

func (s BlestingSnapshot) snapshotToHash() BlestingSnapshot {
	snapshotForHashing := s.Clone()
	for _, fakeResource := range snapshotForHashing.Fakes.List() {
		resources.UpdateMetadata(fakeResource, func(meta *core.Metadata) {
			meta.ResourceVersion = ""
		})
		fakeResource.SetStatus(core.Status{})
	}

	return snapshotForHashing
}

func (s BlestingSnapshot) Hash() uint64 {
	return s.hashStruct(s.snapshotToHash())
}

func (s BlestingSnapshot) HashFields() []zap.Field {
	snapshotForHashing := s.snapshotToHash()
	var fields []zap.Field
	fakes := s.hashStruct(snapshotForHashing.Fakes.List())
	fields = append(fields, zap.Uint64("fakes", fakes))

	return append(fields, zap.Uint64("snapshotHash", s.hashStruct(snapshotForHashing)))
}

func (s BlestingSnapshot) hashStruct(v interface{}) uint64 {
	h, err := hashstructure.Hash(v, nil)
	if err != nil {
		panic(err)
	}
	return h
}
