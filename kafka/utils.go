package kafka

import (
	"gitee.com/eden-framework/common"
	"github.com/segmentio/kafka-go"
)

func wrapKafkaMessage(m kafka.Message) common.QueueMessage {
	return common.QueueMessage{
		Topic:     m.Topic,
		Partition: m.Partition,
		Offset:    m.Offset,
		Key:       m.Key,
		Val:       m.Value,
		Time:      m.Time,
	}
}

func unwrapKafkaMessage(m common.QueueMessage) kafka.Message {
	return kafka.Message{
		Topic:     m.Topic,
		Partition: m.Partition,
		Offset:    m.Offset,
		Key:       m.Key,
		Value:     m.Val,
		Time:      m.Time,
	}
}

func unwrapKafkaMessages(m []common.QueueMessage) (messages []kafka.Message) {
	for _, msg := range m {
		messages = append(messages, unwrapKafkaMessage(msg))
	}
	return
}
