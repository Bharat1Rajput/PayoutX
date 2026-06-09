package outbox

import (
	"context"
	"log"
	"time"
)

func (p *Publisher) Start(
	ctx context.Context,
) {

	ticker := time.NewTicker(
		5 * time.Second,
	)

	defer ticker.Stop()

	for {

		select {

		case <-ctx.Done():
			return

		case <-ticker.C:

			err := p.PublishPending(
				ctx,
			)

			if err != nil {
				log.Printf(
					"outbox publish error: %v",
					err,
				)
			}
		}
	}
}