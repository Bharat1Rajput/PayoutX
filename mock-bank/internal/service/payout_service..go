package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/Bharat1Rajput/payoutX/mock-bank/internal/model"
)

type PayoutService struct{}

func NewPayoutService() *PayoutService {
	return &PayoutService{}
}

func (s *PayoutService) CreatePayout(
	req model.CreatePayoutRequest,
) (*model.CreatePayoutResponse, error) {

	ref := fmt.Sprintf(
		"BANK-%s",
		uuid.NewString(),
	)

	go func() {

		time.Sleep(5 * time.Second)

		webhook := model.BankWebhookRequest{
			PayoutID: req.PayoutID,
			Status:   "SUCCESS",
		}

		body, err := json.Marshal(webhook)
		if err != nil {
			log.Printf("webhook marshal error: %v", err)
			return
		}

		resp, err := http.Post(
			"http://localhost:8080/webhooks/bank",
			"application/json",
			bytes.NewBuffer(body),
		)

		if err != nil {
			log.Printf("webhook send error: %v", err)
			return
		}

		defer resp.Body.Close()

		log.Printf(
			"webhook sent for payout=%s",
			req.PayoutID,
		)

	}()

	return &model.CreatePayoutResponse{
		BankReference: ref,
		Status:        "accepted",
	}, nil
}
