package service

import (
	"context"
	"time"

	"github.com/Bharat1Rajput/payout-service/internal/model"
	"github.com/Bharat1Rajput/payout-service/internal/repository"
	"github.com/google/uuid"
)

type PayoutService struct {
	repo repository.PayoutRepository
}

func NewPayoutService(repo repository.PayoutRepository) *PayoutService {
	return &PayoutService{repo: repo}
}

func (s *PayoutService) CreatePayout(ctx context.Context, req model.CreatePayoutRequest) (*model.CreatePayoutResponse, error) {
	payout := &model.Payout{
		ID:            uuid.NewString(),
		BeneficiaryID: req.BeneficiaryID,
		Amount:        req.Amount,
		Status:        "Created",
		CreatedAt:     time.Now(),
	}

	if err := s.repo.Create(ctx, payout); err != nil {
		return nil, err
	}

	return &model.CreatePayoutResponse{
		ID:     payout.ID,
		Status: payout.Status,
	}, nil

}
