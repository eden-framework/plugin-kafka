package kafka

import (
	"bytes"
	"encoding"
	"errors"

	github_com_eden_framework_eden_framework_pkg_enumeration "github.com/eden-framework/eden-framework/pkg/enumeration"
)

var InvalidBalancerType = errors.New("invalid BalancerType")

func init() {
	github_com_eden_framework_eden_framework_pkg_enumeration.RegisterEnums("BalancerType", map[string]string{
		"MURMUR2":     "murmur2",
		"CRC32":       "crc32",
		"HASH":        "hash",
		"LEAST_BYTES": "least bytes",
		"ROUND_ROBIN": "round robin",
	})
}

func ParseBalancerTypeFromString(s string) (BalancerType, error) {
	switch s {
	case "":
		return BALANCER_TYPE_UNKNOWN, nil
	case "MURMUR2":
		return BALANCER_TYPE__MURMUR2, nil
	case "CRC32":
		return BALANCER_TYPE__CRC32, nil
	case "HASH":
		return BALANCER_TYPE__HASH, nil
	case "LEAST_BYTES":
		return BALANCER_TYPE__LEAST_BYTES, nil
	case "ROUND_ROBIN":
		return BALANCER_TYPE__ROUND_ROBIN, nil
	}
	return BALANCER_TYPE_UNKNOWN, InvalidBalancerType
}

func ParseBalancerTypeFromLabelString(s string) (BalancerType, error) {
	switch s {
	case "":
		return BALANCER_TYPE_UNKNOWN, nil
	case "murmur2":
		return BALANCER_TYPE__MURMUR2, nil
	case "crc32":
		return BALANCER_TYPE__CRC32, nil
	case "hash":
		return BALANCER_TYPE__HASH, nil
	case "least bytes":
		return BALANCER_TYPE__LEAST_BYTES, nil
	case "round robin":
		return BALANCER_TYPE__ROUND_ROBIN, nil
	}
	return BALANCER_TYPE_UNKNOWN, InvalidBalancerType
}

func (BalancerType) EnumType() string {
	return "BalancerType"
}

func (BalancerType) Enums() map[int][]string {
	return map[int][]string{
		int(BALANCER_TYPE__MURMUR2):     {"MURMUR2", "murmur2"},
		int(BALANCER_TYPE__CRC32):       {"CRC32", "crc32"},
		int(BALANCER_TYPE__HASH):        {"HASH", "hash"},
		int(BALANCER_TYPE__LEAST_BYTES): {"LEAST_BYTES", "least bytes"},
		int(BALANCER_TYPE__ROUND_ROBIN): {"ROUND_ROBIN", "round robin"},
	}
}

func (v BalancerType) String() string {
	switch v {
	case BALANCER_TYPE_UNKNOWN:
		return ""
	case BALANCER_TYPE__MURMUR2:
		return "MURMUR2"
	case BALANCER_TYPE__CRC32:
		return "CRC32"
	case BALANCER_TYPE__HASH:
		return "HASH"
	case BALANCER_TYPE__LEAST_BYTES:
		return "LEAST_BYTES"
	case BALANCER_TYPE__ROUND_ROBIN:
		return "ROUND_ROBIN"
	}
	return "UNKNOWN"
}

func (v BalancerType) Label() string {
	switch v {
	case BALANCER_TYPE_UNKNOWN:
		return ""
	case BALANCER_TYPE__MURMUR2:
		return "murmur2"
	case BALANCER_TYPE__CRC32:
		return "crc32"
	case BALANCER_TYPE__HASH:
		return "hash"
	case BALANCER_TYPE__LEAST_BYTES:
		return "least bytes"
	case BALANCER_TYPE__ROUND_ROBIN:
		return "round robin"
	}
	return "UNKNOWN"
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*BalancerType)(nil)

func (v BalancerType) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidBalancerType
	}
	return []byte(str), nil
}

func (v *BalancerType) UnmarshalText(data []byte) (err error) {
	*v, err = ParseBalancerTypeFromString(string(bytes.ToUpper(data)))
	return
}
