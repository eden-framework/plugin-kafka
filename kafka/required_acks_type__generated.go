package kafka

import (
	"bytes"
	"encoding"
	"errors"

	github_com_eden_framework_eden_framework_pkg_enumeration "github.com/eden-framework/eden-framework/pkg/enumeration"
)

var InvalidRequiredAcksType = errors.New("invalid RequiredAcksType")

func init() {
	github_com_eden_framework_eden_framework_pkg_enumeration.RegisterEnums("RequiredAcksType", map[string]string{
		"ALL":  "all",
		"ONE":  "one",
		"NONE": "none",
	})
}

func ParseRequiredAcksTypeFromString(s string) (RequiredAcksType, error) {
	switch s {
	case "":
		return REQUIRED_ACKS_TYPE_UNKNOWN, nil
	case "ALL":
		return REQUIRED_ACKS_TYPE__ALL, nil
	case "ONE":
		return REQUIRED_ACKS_TYPE__ONE, nil
	case "NONE":
		return REQUIRED_ACKS_TYPE__NONE, nil
	}
	return REQUIRED_ACKS_TYPE_UNKNOWN, InvalidRequiredAcksType
}

func ParseRequiredAcksTypeFromLabelString(s string) (RequiredAcksType, error) {
	switch s {
	case "":
		return REQUIRED_ACKS_TYPE_UNKNOWN, nil
	case "all":
		return REQUIRED_ACKS_TYPE__ALL, nil
	case "one":
		return REQUIRED_ACKS_TYPE__ONE, nil
	case "none":
		return REQUIRED_ACKS_TYPE__NONE, nil
	}
	return REQUIRED_ACKS_TYPE_UNKNOWN, InvalidRequiredAcksType
}

func (RequiredAcksType) EnumType() string {
	return "RequiredAcksType"
}

func (RequiredAcksType) Enums() map[int][]string {
	return map[int][]string{
		int(REQUIRED_ACKS_TYPE__ALL):  {"ALL", "all"},
		int(REQUIRED_ACKS_TYPE__ONE):  {"ONE", "one"},
		int(REQUIRED_ACKS_TYPE__NONE): {"NONE", "none"},
	}
}

func (v RequiredAcksType) String() string {
	switch v {
	case REQUIRED_ACKS_TYPE_UNKNOWN:
		return ""
	case REQUIRED_ACKS_TYPE__ALL:
		return "ALL"
	case REQUIRED_ACKS_TYPE__ONE:
		return "ONE"
	case REQUIRED_ACKS_TYPE__NONE:
		return "NONE"
	}
	return "UNKNOWN"
}

func (v RequiredAcksType) Label() string {
	switch v {
	case REQUIRED_ACKS_TYPE_UNKNOWN:
		return ""
	case REQUIRED_ACKS_TYPE__ALL:
		return "all"
	case REQUIRED_ACKS_TYPE__ONE:
		return "one"
	case REQUIRED_ACKS_TYPE__NONE:
		return "none"
	}
	return "UNKNOWN"
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*RequiredAcksType)(nil)

func (v RequiredAcksType) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidRequiredAcksType
	}
	return []byte(str), nil
}

func (v *RequiredAcksType) UnmarshalText(data []byte) (err error) {
	*v, err = ParseRequiredAcksTypeFromString(string(bytes.ToUpper(data)))
	return
}
