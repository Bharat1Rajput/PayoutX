package reconciliation

import (
	"context"
	"log"
	"time"
)

func (s *Service) Start(
	ctx context.Context,
) {

	ticker := time.NewTicker(30 * time.Second)

	defer ticker.Stop()

	for {

		select {

		case <-ctx.Done():
			return

		case <-ticker.C:

			log.Println(
				"starting reconciliation...",
			)

			err := s.Reconcile(ctx)

			if err != nil {
				log.Printf(
					"reconciliation error: %v",
					err,
				)
			}
		}
	}
}