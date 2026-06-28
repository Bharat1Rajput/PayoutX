package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/Bharat1Rajput/payoutX/mock-bank/internal/model"
	"github.com/Bharat1Rajput/payoutX/mock-bank/internal/store"
	"github.com/google/uuid"
)

type PayoutService struct{}

func NewPayoutService() *PayoutService {
	return &PayoutService{}
}
func (s *PayoutService) CreatePayout(
	req model.CreatePayoutRequest,
) (*model.CreatePayoutResponse, error) {
	log.Printf(
		"received payout request: payout_id=%s amount=%d",
		req.PayoutID,
		req.Amount,
	)

	// 70% failure rate for retry testing
	if rand.Intn(100) < 70 {
		log.Printf(
			"bank temporarily unavailable for payout=%s",
			req.PayoutID,
		)

		return nil, errors.New(
			"temporary bank failure",
		)
	}

	ref := fmt.Sprintf(
		"BANK-%s",
		uuid.NewString(),
	)

	log.Printf(
		"bank accepted payout=%s ref=%s",
		req.PayoutID,
		ref,
	)

	store.Save(
		store.Payout{
			PayoutID:      req.PayoutID,
			Status:        "SUCCESS",
			BankReference: ref,
		},
	)

	go func() {

		time.Sleep(5 * time.Second)

		webhook := model.BankWebhookRequest{
			PayoutID: req.PayoutID,
			Status:   "SUCCESS",
		}

		body, err := json.Marshal(webhook)
		if err != nil {
			log.Printf(
				"webhook marshal error: %v",
				err,
			)
			return
		}

		resp, err := http.Post(
			"http://localhost:8080/webhooks/bank",
			"application/json",
			bytes.NewBuffer(body),
		)

		if err != nil {
			log.Printf(
				"webhook send error: %v",
				err,
			)
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

func (s *PayoutService) GetPayout(
	payoutID string,
) (*model.GetPayoutResponse, error) {

	payout, ok := store.Get(
		payoutID,
	)

	if !ok {
		return nil, fmt.Errorf(
			"payout not found",
		)
	}

	return &model.GetPayoutResponse{
		PayoutID: payout.PayoutID,
		Status:   payout.Status,
	}, nil
}
