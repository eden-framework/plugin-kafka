package kafka

import (
	"context"
	"gitee.com/eden-framework/common"
	"github.com/profzone/envconfig"
	"github.com/segmentio/kafka-go"
	"time"
)

type Consumer struct {
	Brokers                []string
	GroupID                string
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
}

func (c *Consumer) SetDefaults() {
	if len(c.Brokers) == 0 {
		panic("kafka.Consumer must set a broker list")
	}
}

func (c *Consumer) Init() {
	c.SetDefaults()
}

func (c *Consumer) Consume(ctx context.Context, topic string, handler func(m common.QueueMessage) error) error {
	reader := c.newReader(topic)
	defer reader.Close()
Run:
	for {
		select {
		case <-ctx.Done():
			break Run
		default:
			m, err := reader.FetchMessage(ctx)
			if err != nil {
				return err
			}

			err = handler(wrapKafkaMessage(m))
			if err != nil {
				continue
			}

			if c.GroupID != "" {
				err = reader.CommitMessages(ctx, m)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (c *Consumer) newReader(topic string) (reader *kafka.Reader) {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:         c.Brokers,
		GroupID:         c.GroupID,
		Topic:           topic,
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
