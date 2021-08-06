package kafka_util

import (
	"context"
	"errors"
	"github.com/igorroncevic/xws2021-nistagram/common/logger"
	"github.com/segmentio/kafka-go"
	"strings"
)

func NewRetryHandler() RetryHandler {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9092"},
		Topic:     RetryTopic,
		Partition: 0,
		MaxWait:   RetryConsumerMaxWait,
	})

	w := NewProducer(RetryTopic)

	l := logger.NewLogger()

	return RetryHandler{ Reader: r, Writer: w, Logger: l }
}

type RetryHandler struct {
	Reader *kafka.Reader
	Writer *KafkaProducer
	Logger *logger.Logger
}

const (
	contextDeadlineExceeded = "context deadline exceeded"
	badRunsLimit            = 5
)

func (handler *RetryHandler) TransferToMainTopic() {
	badRunsFetching := 0
	badRunsTransferring := 0
	badRunsCommitting := 0

	for {
		// TODO This is used to prevent infinite loops. Test if necessary.
		if err := determineBadRunError(badRunsFetching, badRunsTransferring, badRunsCommitting); err != nil {
			break
		}

		m, err := handler.fetchWithTimeout(context.Background())
		if err != nil {
			// Kafka has been trying for too long to find messages and has found none, therefore we can quit looking
			kafkaError := strings.TrimSpace(err.Error())
			if kafkaError == contextDeadlineExceeded { break }

			handler.Logger.ToStdout("Kafka Retry Handler", "Error while fetching message from retry topic", logger.Error)
			badRunsFetching++
			continue
		}

		err = handler.Writer.WriteMessageToRetry(string(m.Value), string(m.Key))
		if err != nil {
			handler.Logger.ToStdout("Kafka Retry Handler", "Error while writing message to main topic", logger.Error)
			badRunsTransferring++
			continue
		}

		err = handler.Reader.CommitMessages(context.Background(), m)
		if err != nil {
			handler.Logger.ToStdout("Kafka Retry Handler", "Failed to commit message", logger.Error)
			badRunsCommitting++
			continue
		}
	}
}

func determineBadRunError(fetchNum int, transferNum int, commitNum int) error {
	if fetchNum > badRunsLimit { return errors.New("failed to fetch too many times") }
	if transferNum > badRunsLimit { return errors.New("failed to transfer too many times") }
	if commitNum > badRunsLimit { return errors.New("failed to commit too many times") }

	return nil
}

func (handler *RetryHandler) fetchWithTimeout(ctx context.Context) (kafka.Message, error) {
	ctx, cancel := context.WithTimeout(ctx, RetryConsumerMaxWait)
	defer cancel()
	return handler.Reader.FetchMessage(ctx)
}