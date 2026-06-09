package outbox

import (
	"context"
	"log"

	"github.com/Bharat1Rajput/payoutX/payout-service/internal/kafka"
	"github.com/Bharat1Rajput/payoutX/payout-service/internal/repository"
)

type Publisher struct {
	repo     repository.OutboxRepository
	producer *kafka.Producer
}

func NewPublisher(
	repo repository.OutboxRepository,
	producer *kafka.Producer,
) *Publisher {

	return &Publisher{
		repo:     repo,
		producer: producer,
	}
}

func (p *Publisher) PublishPending(
	ctx context.Context,
) error {

	events, err := p.repo.GetPending(ctx)
	if err != nil {
		return err
	}

	for _, event := range events {

		err := p.producer.Publish(
			ctx,
			[]byte(event.ID),
			event.Payload,
		)

		if err != nil {
			log.Printf(
				"failed to publish outbox event=%s: %v",
				event.ID,
				err,
			)

			continue
		}

		err = p.repo.MarkSent(
			ctx,
			event.ID,
		)

		if err != nil {
			log.Printf(
				"failed to mark event sent=%s: %v",
				event.ID,
				err,
			)

			continue
		}

		log.Printf(
			"outbox event=%s published",
			event.ID,
		)
	}

	return nil
}
