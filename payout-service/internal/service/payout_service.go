package service

import (
	"context"
	"time"

	"github.com/Bharat1Rajput/payoutX/payout-service/internal/model"
	"github.com/Bharat1Rajput/payoutX/payout-service/internal/repository"
	"github.com/google/uuid"

	grpcClient "github.com/Bharat1Rajput/payoutX/payout-service/internal/grpc"
)

type PayoutService struct {
	repo         repository.PayoutRepository
	ledgerClient *grpcClient.LedgerClient
}

func NewPayoutService(
	repo repository.PayoutRepository,
	ledgerClient *grpcClient.LedgerClient,
) *PayoutService {
	return &PayoutService{repo: repo, ledgerClient: ledgerClient}
}

func (s *PayoutService) CreatePayout(ctx context.Context, req model.CreatePayoutRequest) (*model.CreatePayoutResponse, error) {
	payout := &model.Payout{
		ID:            uuid.NewString(),
		BeneficiaryID: req.BeneficiaryID,
		Amount:        req.Amount,
		Status:        "Created",
		CreatedAt:     time.Now(),
	}
	err := s.ledgerClient.CreateLedgerEntries(
		ctx,
		payout.ID,
		payout.Amount,
	)

	if err != nil {
		return nil, err
	}
	if err := s.repo.Create(ctx, payout); err != nil {
		return nil, err
	}

	return &model.CreatePayoutResponse{
		ID:     payout.ID,
		Status: payout.Status,
	}, nil

}
