package kafka

//go:generate eden generate enum --type-name=RequiredAcksType
// api:enum
type RequiredAcksType uint8

// required acks type
const (
	REQUIRED_ACKS_TYPE_UNKNOWN RequiredAcksType = iota
	REQUIRED_ACKS_TYPE__NONE                    // none
	REQUIRED_ACKS_TYPE__ONE                     // one
	REQUIRED_ACKS_TYPE__ALL                     // all
)
