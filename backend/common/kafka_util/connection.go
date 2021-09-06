package kafka_util

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
)

func GetConnection(topic string) (*kafka.Conn, error) {
	partition := 0
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
		return nil, err
	}

	return conn, nil
}
