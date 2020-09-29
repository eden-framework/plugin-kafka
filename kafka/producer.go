package kafka

import (
	"context"
	"fmt"
	"github.com/profzone/envconfig"
	"github.com/segmentio/kafka-go"
	"net"
	"time"
)

type Producer struct {
	Host         string
	Port         int
	Topic        string
	BalancerType BalancerType
	MaxAttempts  int
	BatchSize    int
	BatchBytes   int64
	BatchTimeout envconfig.Duration
	ReadTimeout  envconfig.Duration
	WriteTimeout envconfig.Duration
	RequiredAcks RequiredAcksType
	Async        bool
	w            *kafka.Writer
}

func (p *Producer) SetDefaults() {
	if p.Host == "" {
		panic("kafka.Producer must set a broker Host")
	}
	if p.Topic == "" {
		panic("kafka.Producer must set a topic")
	}

	if p.Port == 0 {
		p.Port = 9092
	}
}

func (p *Producer) Init() {
	if p.w != nil {
		return
	}

	p.SetDefaults()

	addr := fmt.Sprintf("%s:%d", p.Host, p.Port)
	resolvedAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		panic(fmt.Sprintf("kafka address [%s] resolve failed: %v", addr, err))
	}
	p.w = &kafka.Writer{
		Addr:         resolvedAddr,
		Topic:        p.Topic,
		Balancer:     newBalancer(p.BalancerType),
		MaxAttempts:  p.MaxAttempts,
		BatchSize:    p.BatchSize,
		BatchBytes:   p.BatchBytes,
		BatchTimeout: time.Duration(p.BatchTimeout),
		ReadTimeout:  time.Duration(p.ReadTimeout),
		WriteTimeout: time.Duration(p.WriteTimeout),
		RequiredAcks: newRequiredAcks(p.RequiredAcks),
		Async:        p.Async,
	}
}

func (p *Producer) Produce(ctx context.Context, messages ...kafka.Message) error {
	return p.w.WriteMessages(ctx, messages...)
}

func newBalancer(t BalancerType) kafka.Balancer {
	switch t {
	case BALANCER_TYPE__ROUND_ROBIN:
		return &kafka.RoundRobin{}
	case BALANCER_TYPE__LEAST_BYTES:
		return &kafka.LeastBytes{}
	case BALANCER_TYPE__CRC32:
		return &kafka.CRC32Balancer{}
	case BALANCER_TYPE__HASH:
		return &kafka.Hash{}
	case BALANCER_TYPE__MURMUR2:
		return &kafka.Murmur2Balancer{}
	default:
		return &kafka.RoundRobin{}
	}
}

func newRequiredAcks(t RequiredAcksType) kafka.RequiredAcks {
	switch t {
	case REQUIRED_ACKS_TYPE__NONE:
		return kafka.RequireNone
	case REQUIRED_ACKS_TYPE__ONE:
		return kafka.RequireOne
	case REQUIRED_ACKS_TYPE__ALL:
		return kafka.RequireAll
	default:
		return kafka.RequireAll
	}
}
