package worker

import (
	"log"

	"github.com/Bharat1Rajput/payoutX/processor-service/internal/bank"
	"github.com/Bharat1Rajput/payoutX/processor-service/internal/model"
	"github.com/Bharat1Rajput/payoutX/processor-service/internal/payout"
)

type PayoutWorker struct {
	bankClient   *bank.Client
	payoutClient *payout.Client
}

func NewPayoutWorker(bankClient *bank.Client, payoutClient *payout.Client) *PayoutWorker {
	return &PayoutWorker{
		bankClient:   bankClient,
		payoutClient: payoutClient,
	}
}

func (w *PayoutWorker) ExecutePayout(
	event model.PayoutCreatedEvent,
) error {


	log.Println("processing payout")
	err := w.payoutClient.UpdateStatus(
		event.PayoutID,
		"PROCESSING",
	)

	if err != nil {
		return err
	}

	log.Printf(
		"payout=%s marked as processing",
		event.PayoutID,
	)

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
