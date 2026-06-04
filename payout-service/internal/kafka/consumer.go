package kafka

import (
	"context"
	"log"

	kafkago "github.com/segmentio/kafka-go"
)

type Consumer struct {
	reader *kafkago.Reader
}

func NewConsumer(
	broker string,
	topic string,
	groupID string,
) *Consumer {

	reader := kafkago.NewReader(kafkago.ReaderConfig{
		Brokers: []string{broker},
		Topic:   topic,
		GroupID: groupID,
	})

	return &Consumer{
		reader: reader,
	}
}

func (c *Consumer) Start(ctx context.Context) {
	for {
		msg, err := c.reader.ReadMessage(ctx)
		if err != nil {
			log.Printf("error reading message: %v", err)
			continue
		}

		log.Printf(
			"received event: key=%s value=%s",
			string(msg.Key),
			string(msg.Value),
		)
	}
}