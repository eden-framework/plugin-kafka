package kafka

import (
	"context"
	"github.com/profzone/envconfig"
	"github.com/segmentio/kafka-go"
	"time"
)

type Consumer struct {
	Brokers                []string
	GroupID                string
	Topic                  string
	Partition              int
	QueueCapacity          int
	MinBytes               int
	MaxBytes               int
	MaxWait                envconfig.Duration
	ReadLagInterval        envconfig.Duration
	HeartbeatInterval      envconfig.Duration
	CommitInterval         envconfig.Duration
	PartitionWatchInterval envconfig.Duration
	WatchPartitionChanges  bool
	SessionTimeout         envconfig.Duration
	RebalanceTimeout       envconfig.Duration
	JoinGroupBackoff       envconfig.Duration
	RetentionTime          envconfig.Duration
	StartOffset            int64
	ReadBackoffMin         envconfig.Duration
	ReadBackoffMax         envconfig.Duration
	MaxAttempts            int

	r *kafka.Reader
}

func (c *Consumer) SetDefaults() {
	if len(c.Brokers) == 0 {
		panic("kafka.Consumer must set a broker list")
	}
	if c.Topic == "" {
		panic("kafka.Consumer must set a topic")
	}
}

func (c *Consumer) Init() {
	c.r = kafka.NewReader(kafka.ReaderConfig{
		Brokers:         c.Brokers,
		GroupID:         c.GroupID,
		Topic:           c.Topic,
		Partition:       c.Partition,
		QueueCapacity:   c.QueueCapacity,
		MinBytes:        c.MinBytes,
		MaxBytes:        c.MaxBytes,
		MaxWait:         time.Duration(c.MaxWait),
		ReadLagInterval: time.Duration(c.ReadLagInterval),
		//GroupBalancers:         nil,
		HeartbeatInterval:      time.Duration(c.HeartbeatInterval),
		CommitInterval:         time.Duration(c.CommitInterval),
		PartitionWatchInterval: time.Duration(c.PartitionWatchInterval),
		WatchPartitionChanges:  c.WatchPartitionChanges,
		SessionTimeout:         time.Duration(c.SessionTimeout),
		RebalanceTimeout:       time.Duration(c.RebalanceTimeout),
		JoinGroupBackoff:       time.Duration(c.JoinGroupBackoff),
		RetentionTime:          time.Duration(c.RetentionTime),
		StartOffset:            c.StartOffset,
		ReadBackoffMin:         time.Duration(c.ReadBackoffMin),
		ReadBackoffMax:         time.Duration(c.ReadBackoffMax),
		MaxAttempts:            c.MaxAttempts,
	})
}

func (c *Consumer) ReadMessage(ctx context.Context, handler func(m kafka.Message) error) error {
Run:
	for {
		select {
		case <-ctx.Done():
			break Run
		default:
			m, err := c.r.FetchMessage(ctx)
			if err != nil {
				return err
			}

			err = handler(m)
			if err != nil {
				continue
			}

			if c.GroupID != "" {
				err = c.r.CommitMessages(ctx, m)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
