package service

import (
	"fmt"

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

	return &model.CreatePayoutResponse{
		BankReference: ref,
		Status:        "accepted",
	}, nil
}