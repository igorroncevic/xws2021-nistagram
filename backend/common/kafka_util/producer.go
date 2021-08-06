package kafka_util

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"github.com/segmentio/kafka-go"
	"log"
)

func NewProducer(topic string) *KafkaProducer {
	writer := &kafka.Writer{
		Addr:         kafka.TCP("localhost:9092"),
		Topic:        topic,
		Balancer:     &kafka.LeastBytes{},
	}

	return &KafkaProducer{ writer }
}

type KafkaProducer struct {
	Writer *kafka.Writer
}

func (producer *KafkaProducer) WriteMessage(message string) error {
	kafkaMessage := kafka.Message{
		Key:   []byte(uuid.NewV4().String()),
		Value: []byte(message),
	}

	err := producer.Writer.WriteMessages(context.Background(), kafkaMessage)

	if err != nil {
		log.Println("failed to write messages to '" + producer.Writer.Topic + "' topic: ", err.Error())
		return err
	}

	if err := producer.Writer.Close(); err != nil {
		log.Println("failed to close writer: ", err)
		return err
	}

	return err
}

func (producer *KafkaProducer) WriteMessageToRetry(message string, key string) error {
	kafkaMessage := kafka.Message{
		Key:   []byte(key),
		Value: []byte(message),
	}

	err := producer.Writer.WriteMessages(context.Background(), kafkaMessage)

	if err != nil {
		log.Println("failed to write messages to '" + producer.Writer.Topic + "' topic: ", err.Error())
		return err
	}

	if err := producer.Writer.Close(); err != nil {
		log.Println("failed to close writer: ", err)
		return err
	}

	return err
}

