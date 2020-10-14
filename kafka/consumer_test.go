package kafka

import (
	"context"
	"fmt"
	"github.com/eden-framework/common"
	"testing"
)

func TestConsume(t *testing.T) {
	c := &Consumer{
		Brokers:     []string{"localhost:9092"},
		GroupID:     "group1",
		MinBytes:    1e3,
		MaxBytes:    1e6,
		StartOffset: 0,
	}
	c.Init()

	c.Consume(context.Background(), "test-A", func(m common.QueueMessage) error {
		fmt.Printf("message of offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Val))
		return nil
	})
}

func TestTwoConsumer(t *testing.T) {
	c1 := &Consumer{
		Brokers: []string{"localhost:9092"},
		GroupID: "group1",
	}
	c1.Init()

	c2 := &Consumer{
		Brokers: []string{"localhost:9092"},
		GroupID: "group1",
	}
	c2.Init()

	go func() {
		var receivedNum uint32
		err := c1.Consume(context.Background(), "test-B", func(m common.QueueMessage) error {
			receivedNum++
			fmt.Printf("[%d] message of offset %d: %s = %s\n", receivedNum, m.Offset, string(m.Key), string(m.Val))
			return nil
		})
		if err != nil {
			return
		}
	}()

	go func() {
		var receivedNum uint32
		err := c2.Consume(context.Background(), "test-B", func(m common.QueueMessage) error {
			receivedNum++
			fmt.Printf("\t[%d] message of offset %d: %s = %s\n", receivedNum, m.Offset, string(m.Key), string(m.Val))
			return nil
		})
		if err != nil {
			return
		}
	}()

	select {}
}
