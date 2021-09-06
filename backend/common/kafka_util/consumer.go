package kafka_util

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/igorroncevic/xws2021-nistagram/common/logger"
	"github.com/segmentio/kafka-go"
	"os"
	"time"
)

func NewConsumer(maxWait time.Duration, groupId string, topicName string) KafkaConsumer {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{os.Getenv("KAFKA_HOST") + ":9092"},
		Topic:     topicName,
		Partition: 0,
		MaxWait:   maxWait,
		GroupID:   groupId,
	})
	r.SetOffset(kafka.LastOffset)

	l := logger.NewLogger()

	l.ToStdout("Kafka Consumer", "Ready on topic "+topicName, logger.Info)

	return KafkaConsumer{reader: r, logger: l}
}

type KafkaConsumer struct {
	reader *kafka.Reader
	logger *logger.Logger
}

func (consumer *KafkaConsumer) Consume(consumeLogic func(string, time.Time, map[string]interface{}) error) {
	retryProducer := NewProducer(RetryTopic)

	for {
		// Try to read the message
		message, err := consumer.reader.ReadMessage(context.Background())
		if err != nil {
			consumer.logger.ToStdout("Kafka Consumer", "Error while fetching message", logger.Error)
			consumer.resolveError(retryProducer, message)
			continue
		}

		// Log the incoming message
		consumer.logger.ToStdout("Kafka Consumer",
			fmt.Sprintf("message at offset %d: %s = %s\n", message.Offset, string(message.Key), string(message.Value)),
			logger.Info)

		// Parse the message
		var jsonMessage map[string]interface{}
		err = json.Unmarshal(message.Value, &jsonMessage)
		if err != nil {
			consumer.logger.ToStdout("Kafka Consumer", "Unable to parse incoming message", logger.Error)
			consumer.resolveError(retryProducer, message)
			continue
		}

		// Consume logic
		if consumeLogic == nil {
			continue
		}
		err = consumeLogic(string(message.Key), message.Time, jsonMessage)
		if err != nil {
			consumer.logger.ToStdout("Kafka Consumer", "Failed to perform consume logic", logger.Error)
			continue
		}

		// Commit the message <=> acknowledge that the message has been consumed
		err = consumer.reader.CommitMessages(context.Background(), message)
		if err != nil {
			// TODO Check if the uncommitted message is read twice
			consumer.logger.ToStdout("Kafka Consumer", "Failed to commit message", logger.Error)
			continue
		}
		consumer.logger.ToStdout("Kafka Consumer", "Committed message "+string(message.Key), logger.Info)
	}

}

// resolveError is used to store a message in retry topic if an error occurs
// if the current topic is a retry topic, commit the message as we are unable to process it
func (consumer *KafkaConsumer) resolveError(producer *KafkaProducer, message kafka.Message) {
	if consumer.reader.Stats().Topic != RetryTopic {
		consumer.writeToRetry(producer, message)
	} else {
		err := consumer.reader.CommitMessages(context.Background(), message)
		if err != nil {
			consumer.logger.ToStdout("Kafka Consumer", "Failed to commit message", logger.Error)
		}
	}
}

func (consumer *KafkaConsumer) writeToRetry(producer *KafkaProducer, message kafka.Message) {
	err := producer.WriteMessageToRetry(string(message.Value), string(message.Key))
	if err != nil {
		consumer.logger.ToStdout("Kafka Consumer", "Failed to write message to retry topic", logger.Error)
	} else {
		consumer.logger.ToStdout("Kafka Consumer", "Successfully written to retry topic", logger.Info)
	}
}

func (consumer *KafkaConsumer) Close() error {
	return consumer.reader.Close()
}
