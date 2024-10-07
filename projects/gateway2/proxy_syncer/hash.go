package proxy_syncer

import (
	"encoding/binary"
	"fmt"
	"hash"
	"hash/fnv"
	"io"
	"math"

	envoy_config_core_v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	structpb "google.golang.org/protobuf/types/known/structpb"
)

func HashMetadata(newhash func() hash.Hash64, md *envoy_config_core_v3.Metadata) uint64 {
	var finalHash uint64
	// Iterate over the sorted keys and add them and their values to the hash
	for key, value := range md.FilterMetadata {
		h := newhash()
		// Write the field name (key) to the hash
		_, err := h.Write([]byte(key))
		if err != nil {
			panic(fmt.Errorf("failed to hash key: %w", err))
		}
		hashUint64(h, HashProtoStruct(newhash, value))
		finalHash ^= h.Sum64()
	}

	return finalHash
}

// HashProtoStruct hashes a protobuf Struct using the provided hash.Hash interface
func HashProtoStruct(newhash func() hash.Hash64, pbStruct *structpb.Struct) uint64 {
	var finalHash uint64
	// Iterate over the sorted keys and add them and their values to the hash
	for key, value := range pbStruct.Fields {
		h := newhash()
		// Write the field name (key) to the hash
		_, err := h.Write([]byte(key))
		if err != nil {
			panic(fmt.Errorf("failed to hash key: %w", err))
		}

		// Write the value to the hash
		if err := hashValue(newhash, h, value); err != nil {
			panic(fmt.Errorf("failed to hash value: %w", err))
		}
		finalHash ^= h.Sum64()
	}

	return finalHash
}

// hashValue recursively hashes a protobuf Value based on its type
func hashValue(newhash func() hash.Hash64, h hash.Hash, value *structpb.Value) error {
	switch kind := value.Kind.(type) {
	case *structpb.Value_NullValue:
		// Null values can just be hashed as a constant
		_, err := h.Write([]byte("null"))
		return err
	case *structpb.Value_NumberValue:
		// Hash the number value
		hashUint64(h, math.Float64bits(kind.NumberValue))
	case *structpb.Value_StringValue:
		// Hash the string value
		_, err := h.Write([]byte(kind.StringValue))
		return err
	case *structpb.Value_BoolValue:
		// Hash the boolean value
		boolStr := "false"
		if kind.BoolValue {
			boolStr = "true"
		}
		_, err := h.Write([]byte(boolStr))
		return err
	case *structpb.Value_StructValue:
		// Recursively hash the struct value
		hashUint64(h, HashProtoStruct(newhash, kind.StructValue))
	case *structpb.Value_ListValue:
		// Recursively hash each element of the list
		for _, elem := range kind.ListValue.Values {
			if err := hashValue(newhash, h, elem); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unknown kind in protobuf Value: %T", kind)
	}
	return nil
}

func hashUint64(hasher io.Writer, value uint64) {
	var bytes [8]byte
	binary.NativeEndian.PutUint64(bytes[:], value)
	hasher.Write(bytes[:])
}

func hashLabels(labels map[string]string) uint64 {
	finalHash := uint64(0)
	for k, v := range labels {
		fnv := fnv.New64()
		fnv.Write([]byte(k))
		fnv.Write([]byte{0})
		fnv.Write([]byte(v))
		fnv.Write([]byte{0})
		finalHash ^= fnv.Sum64()
	}
	return finalHash
}