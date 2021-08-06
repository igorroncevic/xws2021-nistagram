package kafka_util

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/igorroncevic/xws2021-nistagram/common/logger"
	"github.com/segmentio/kafka-go"
	"time"
)

func NewConsumer(maxWait time.Duration, topicName string) KafkaConsumer {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9092"},
		Topic:     topicName,
		Partition: 0,
		MaxWait:   maxWait,
	})
	r.SetOffset(kafka.LastOffset)

	l := logger.NewLogger()

	return KafkaConsumer{ Reader: r, Logger: l }
}

type KafkaConsumer struct {
	Reader *kafka.Reader
	Logger *logger.Logger
}

func (consumer *KafkaConsumer) Consume() {
	consumer.Logger.ToStdout("Kafka Consumer", "Ready on topic " + consumer.Reader.Stats().Topic, logger.Info)

	retryProducer := NewProducer(RetryTopic)
	defer retryProducer.Writer.Close()

	go func() {
		for {
			// Try to read the message
			message, err := consumer.Reader.ReadMessage(context.Background())
			if err != nil {
				consumer.Logger.ToStdout("Kafka Consumer", "Error while fetching message", logger.Error)
				consumer.resolveError(retryProducer, message)
				continue
			}

			// Log the incoming message
			consumer.Logger.ToStdout("Kafka Consumer",
				fmt.Sprintf("message at offset %d: %s = %s\n", message.Offset, string(message.Key), string(message.Value)),
				logger.Info)

			// Parse the message
			var jsonMessage map[string]interface{}
			err = json.Unmarshal(message.Value, &jsonMessage)
			if err != nil {
				consumer.Logger.ToStdout("Kafka Consumer", "Unable to parse incoming message", logger.Error)
				consumer.resolveError(retryProducer, message)
				continue
			}

			// Consummation logic
			// ...
			// End of consummation logic

			// Commit the message <=> acknowledge that the message has been consumed
			err = consumer.Reader.CommitMessages(context.Background(), message)
			if err != nil {
				// TODO Check if the uncommitted message is read twice
				consumer.Logger.ToStdout("Kafka Consumer", "Failed to commit message", logger.Error)
				continue
			}
		}
	}()
}

// resolveError is used to store a message in retry topic if an error occurs
// if the current topic is a retry topic, commit the message as we are unable to process it
func (consumer *KafkaConsumer) resolveError(producer *KafkaProducer, message kafka.Message) {
	if consumer.Reader.Stats().Topic != RetryTopic {
		consumer.writeToRetry(producer, message)
	}else{
		err := consumer.Reader.CommitMessages(context.Background(), message)
		if err != nil {
			consumer.Logger.ToStdout("Kafka Consumer", "Failed to commit message", logger.Error)
		}
	}
}

func  (consumer *KafkaConsumer) writeToRetry(producer *KafkaProducer, message kafka.Message) {
	err := producer.WriteMessageToRetry(string(message.Value), string(message.Key))
	if err != nil {
		consumer.Logger.ToStdout("Kafka Consumer", "Failed to write message to retry topic", logger.Error)
	} else {
		consumer.Logger.ToStdout("Kafka Consumer", "Successfully written to retry topic", logger.Info)
	}
}

