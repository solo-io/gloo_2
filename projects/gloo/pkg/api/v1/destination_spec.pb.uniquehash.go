// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/destination_spec.proto

package v1

import (
	"encoding/binary"
	"errors"
	"fmt"
	"hash"
	"hash/fnv"
	"strconv"

	safe_hasher "github.com/solo-io/protoc-gen-ext/pkg/hasher"
	"github.com/solo-io/protoc-gen-ext/pkg/hasher/hashstructure"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = new(hash.Hash64)
	_ = fnv.New64
	_ = strconv.Itoa
	_ = hashstructure.Hash
	_ = new(safe_hasher.SafeHasher)
)

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *DestinationSpec) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gloo.solo.io.github.com/solo-io/gloo/projects/gloo/pkg/api/v1.DestinationSpec")); err != nil {
		return 0, err
	}

	switch m.DestinationType.(type) {

	case *DestinationSpec_Aws:

		if h, ok := interface{}(m.GetAws()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("Aws")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetAws(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("Aws")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *DestinationSpec_Azure:

		if h, ok := interface{}(m.GetAzure()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("Azure")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetAzure(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("Azure")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *DestinationSpec_Rest:

		if h, ok := interface{}(m.GetRest()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("Rest")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetRest(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("Rest")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *DestinationSpec_Grpc:

		if h, ok := interface{}(m.GetGrpc()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("Grpc")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetGrpc(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("Grpc")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	return hasher.Sum64(), nil
}
