package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/Bharat1Rajput/payoutX/processor-service/internal/bank"
	"github.com/Bharat1Rajput/payoutX/processor-service/internal/model"
	"github.com/Bharat1Rajput/payoutX/processor-service/internal/payout"
	"github.com/Bharat1Rajput/payoutX/processor-service/internal/port"
)

type PayoutWorker struct {
	bankClient   *bank.Client
	payoutClient *payout.Client
	dlqProducer  port.DLQPublisher
}

func NewPayoutWorker(
	bankClient *bank.Client,
	payoutClient *payout.Client,
	dlqProducer port.DLQPublisher,
) *PayoutWorker {

	return &PayoutWorker{
		bankClient:   bankClient,
		payoutClient: payoutClient,
		dlqProducer:  dlqProducer,
	}
}

func (w *PayoutWorker) ExecutePayout(
	event model.PayoutCreatedEvent,
) error {

	log.Printf(
		"starting payout processing: payout_id=%s",
		event.PayoutID,
	)

	err := w.payoutClient.UpdateStatus(
		event.PayoutID,
		model.PayoutProcessing,
	)

	if err != nil {
		return err
	}

	log.Printf(
		"payout=%s marked as %s",
		event.PayoutID,
		model.PayoutProcessing,
	)

	var (
		resp    *bank.CreatePayoutResponse
		lastErr error
	)

	const maxRetries = 3

	for attempt := 1; attempt <= maxRetries; attempt++ {

		log.Printf(
			"calling bank: payout=%s attempt=%d/%d",
			event.PayoutID,
			attempt,
			maxRetries,
		)
		resp, lastErr = w.bankClient.CreatePayout(
			bank.CreatePayoutRequest{
				PayoutID: event.PayoutID,
				Amount:   event.Amount,
			},
		)

		log.Printf(
			"bank response err=%v",
			lastErr,
		)

		if lastErr == nil {
			break
		}

		log.Printf(
			"bank call failed: payout=%s attempt=%d error=%v",
			event.PayoutID,
			attempt,
			lastErr,
		)

		if attempt < maxRetries {
			time.Sleep(2 * time.Second)
		}
	}

	if lastErr != nil {

		// Mark payout as FAILED
		err := w.payoutClient.UpdateStatus(
			event.PayoutID,
			model.PayoutFailed,
		)

		if err != nil {
			log.Printf(
				"failed to mark payout as failed: %v",
				err,
			)
		}

		// Publish failed event to DLQ
		failedEvent := model.PayoutFailedEvent{
			PayoutID: event.PayoutID,
			Amount:   event.Amount,
			Reason:   lastErr.Error(),
		}

		eventBytes, err := json.Marshal(
			failedEvent,
		)

		if err != nil {
			return err
		}

		err = w.dlqProducer.Publish(
			context.Background(),
			[]byte(event.PayoutID),
			eventBytes,
		)

		if err != nil {
			log.Printf(
				"failed to publish dlq event: %v",
				err,
			)
		}

		log.Printf(
			"payout=%s sent to dlq",
			event.PayoutID,
		)

		return fmt.Errorf(
			"bank call failed after %d attempts: %w",
			maxRetries,
			lastErr,
		)
	}

	log.Printf(
		"bank accepted payout=%s ref=%s",
		event.PayoutID,
		resp.BankReference,
	)

	return nil
}
