package kafka

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"testing"
)

func TestConsume(t *testing.T) {
	c := &Consumer{
		Brokers:     []string{"localhost:9092"},
		Topic:       "test-A",
		GroupID:     "group1",
		MinBytes:    1e3,
		MaxBytes:    1e6,
		StartOffset: 0,
	}
	c.Init()

	for {
		m, err := c.r.ReadMessage(context.Background())
		if err != nil {
			break
		}

		fmt.Printf("message of offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	}
}

func TestTwoConsumer(t *testing.T) {
	c1 := &Consumer{
		Brokers: []string{"localhost:9092"},
		Topic:   "test-B",
		GroupID: "group1",
	}
	c1.Init()

	c2 := &Consumer{
		Brokers: []string{"localhost:9092"},
		Topic:   "test-B",
		GroupID: "group1",
	}
	c2.Init()

	go func() {
		var receivedNum uint32
		err := c1.ReadMessage(context.Background(), func(m kafka.Message) error {
			receivedNum++
			fmt.Printf("[%d] message of offset %d: %s = %s\n", receivedNum, m.Offset, string(m.Key), string(m.Value))
			return nil
		})
		if err != nil {
			return
		}
	}()

	go func() {
		var receivedNum uint32
		err := c2.ReadMessage(context.Background(), func(m kafka.Message) error {
			receivedNum++
			fmt.Printf("\t[%d] message of offset %d: %s = %s\n", receivedNum, m.Offset, string(m.Key), string(m.Value))
			return nil
		})
		if err != nil {
			return
		}
	}()

	select {}
}
