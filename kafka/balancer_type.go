package kafka

//go:generate eden generate enum --type-name=BalancerType
// api:enum
type BalancerType uint8

// kafka balancer type
const (
	BALANCER_TYPE_UNKNOWN      BalancerType = iota
	BALANCER_TYPE__ROUND_ROBIN              // round robin
	BALANCER_TYPE__LEAST_BYTES              // least bytes
	BALANCER_TYPE__HASH                     // hash
	BALANCER_TYPE__CRC32                    // crc32
	BALANCER_TYPE__MURMUR2                  // murmur2
)
