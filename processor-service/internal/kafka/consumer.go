package kafka

import (
	"context"
	"encoding/json"
	"log"

	"github.com/Bharat1Rajput/payoutX/processor-service/internal/model"
	"github.com/Bharat1Rajput/payoutX/processor-service/internal/worker"
	kafkago "github.com/segmentio/kafka-go"
)

type Consumer struct {
	reader *kafkago.Reader
	worker *worker.PayoutWorker
}

func NewConsumer(
	broker string,
	topic string,
	groupID string,
	worker *worker.PayoutWorker,
) *Consumer {

	reader := kafkago.NewReader(kafkago.ReaderConfig{
		Brokers: []string{broker},
		Topic:   topic,
		GroupID: groupID,
	})

	return &Consumer{
		reader: reader,
		worker: worker,
	}
}

func (c *Consumer) Start(ctx context.Context) {

	for {
		msg, err := c.reader.ReadMessage(ctx)
		if err != nil {
			log.Printf("error reading message: %v", err)
			continue
		}

		var event model.PayoutCreatedEvent

		err = json.Unmarshal(msg.Value, &event)
		if err != nil {
			log.Printf("error unmarshalling message: %v", err)
			continue
		}

		err = c.worker.ExecutePayout(event)
		if err != nil {
			log.Printf("error executing payout: %v", err)
			continue
		}
	}
}
