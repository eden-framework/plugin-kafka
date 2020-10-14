package kafka

import (
	"context"
	"fmt"
	"github.com/eden-framework/common"
	"github.com/gophercloud/gophercloud/acceptance/tools"
	"github.com/segmentio/kafka-go"
	"testing"
)

func TestCreateTopic(t *testing.T) {
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "default", 0)
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()
	topicConfigs := []kafka.TopicConfig{
		{
			Topic:             "default",
			NumPartitions:     3,
			ReplicationFactor: 3,
		},
	}

	err = conn.CreateTopics(topicConfigs...)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteTopic(t *testing.T) {

}

func TestProduce(t *testing.T) {
	p := &Producer{
		Host: "127.0.0.1",
		Port: 9092,
	}
	p.Init()

	messages := make([]common.QueueMessage, 0)
	for i := 0; i < 1000; i++ {
		messages = append(messages, common.QueueMessage{
			Topic: "test-B",
			Key:   []byte(tools.RandomString("", 10)),
			Val:   []byte(fmt.Sprintf("bar%d", i+1)),
		})
	}
	err := p.Produce(context.Background(), messages...)

	if err != nil {
		t.Fatal(err)
	}
}
