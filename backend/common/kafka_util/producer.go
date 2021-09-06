package kafka_util

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"github.com/segmentio/kafka-go"
	"log"
	"os"
	"time"
)

func NewProducer(topic string) *KafkaProducer {
	writer := &kafka.Writer{
		Addr:     kafka.TCP(os.Getenv("KAFKA_HOST") + ":9092"),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}

	return &KafkaProducer{writer}
}

type KafkaProducer struct {
	Writer *kafka.Writer
}

func (producer *KafkaProducer) WriteMessage(message string) error {
	kafkaMessage := kafka.Message{
		Key:   []byte(uuid.NewV4().String()),
		Value: []byte(message),
		Time:  time.Now(),
	}

	err := producer.Writer.WriteMessages(context.Background(), kafkaMessage)

	if err != nil {
		log.Println("failed to write messages to '"+producer.Writer.Topic+"' topic: ", err.Error())
		producer.WriteMessageToRetry(message, string(kafkaMessage.Key))
		return err
	}

	return err
}

func (producer *KafkaProducer) WriteMessageToRetry(message string, key string) error {
	kafkaMessage := kafka.Message{
		Key:   []byte(key),
		Value: []byte(message),
		Time:  time.Now(),
	}

	err := producer.Writer.WriteMessages(context.Background(), kafkaMessage)

	if err != nil {
		log.Println("failed to write messages to '"+producer.Writer.Topic+"' topic: ", err.Error())
		return err
	}

	return err
}

func (producer *KafkaProducer) WritePerformanceMessage(service string, function string, message string, status int) error {
	message, err := MarshalPerformanceMessage(service, function, status, message)
	if err != nil {
		log.Println("Couldn't marshal Kafka message")
		//s.logger.ToStdoutAndFile("KafkaCallback", "Couldn't marshal Kafka message", logger.Error)
		return err
	}

	err = producer.WriteMessage(message)
	if err != nil {
		log.Println("Couldn't log event to Kafka")
		//s.logger.ToStdoutAndFile("KafkaCallback", "Couldn't log event to Kafka", logger.Error)
		return err
	}

	return nil
}

func (producer *KafkaProducer) WriteUserEventMessage(eventType UserEventType, userId string, message string) error {
	message, err := MarshalUserEventMessage(eventType, userId, message)
	if err != nil {
		log.Println("Couldn't marshal Kafka message")
		//s.logger.ToStdoutAndFile("KafkaCallback", "Couldn't marshal Kafka message", logger.Error)
		return err
	}

	err = producer.WriteMessage(message)
	if err != nil {
		log.Println("Couldn't log event to Kafka")
		//s.logger.ToStdoutAndFile("KafkaCallback", "Couldn't log event to Kafka", logger.Error)
		return err
	}

	return nil
}
