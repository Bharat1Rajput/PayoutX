package service

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Bharat1Rajput/payoutX/payout-service/internal/kafka"
	"github.com/Bharat1Rajput/payoutX/payout-service/internal/model"
	"github.com/Bharat1Rajput/payoutX/payout-service/internal/repository"
	"github.com/google/uuid"

	grpcClient "github.com/Bharat1Rajput/payoutX/payout-service/internal/grpc"
)

type PayoutService struct {
	repo          repository.PayoutRepository
	ledgerClient  *grpcClient.LedgerClient
	kafkaProducer *kafka.Producer
}

func NewPayoutService(
	repo repository.PayoutRepository,
	ledgerClient *grpcClient.LedgerClient,
	kafkaProducer *kafka.Producer,
) *PayoutService {
	return &PayoutService{repo: repo, ledgerClient: ledgerClient, kafkaProducer: kafkaProducer}
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

	event := kafka.PayoutCreatedEvent{
		PayoutID:      payout.ID,
		BeneficiaryID: payout.BeneficiaryID,
		Amount:        payout.Amount,
		Status:        payout.Status,
	}

	eventBytes, err := json.Marshal(event)
	if err != nil {
		return nil, err
	}

	err = s.kafkaProducer.Publish(
		ctx,
		[]byte(payout.ID),
		eventBytes,
	)

	if err != nil {
		return nil, err
	}

	return &model.CreatePayoutResponse{
		ID:     payout.ID,
		Status: payout.Status,
	}, nil

}

func (s *PayoutService) UpdatePayoutStatus(
	ctx context.Context,
	payoutID string,
	status string,
) error {

	return s.repo.UpdateStatus(
		ctx,
		payoutID,
		status,
	)
}

