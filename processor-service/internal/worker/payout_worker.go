package worker

import (
	"log"

	"github.com/Bharat1Rajput/payoutX/processor-service/internal/bank"
	"github.com/Bharat1Rajput/payoutX/processor-service/internal/model"
)

type PayoutWorker struct {
	bankClient *bank.Client
}

func NewPayoutWorker(bankClient *bank.Client) *PayoutWorker {
	return &PayoutWorker{
		bankClient: bankClient,
	}
}

func (w *PayoutWorker) ExecutePayout(
	event model.PayoutCreatedEvent,
) error {

	resp, err := w.bankClient.CreatePayout(
		bank.CreatePayoutRequest{
			PayoutID: event.PayoutID,
			Amount:   event.Amount,
		},
	)

	if err != nil {
		return err
	}

	log.Printf(
		"bank accepted payout=%s ref=%s",
		event.PayoutID,
		resp.BankReference,
	)

	return nil
}