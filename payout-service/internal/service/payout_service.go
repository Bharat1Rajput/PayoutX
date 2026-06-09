package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/Bharat1Rajput/payoutX/payout-service/internal/kafka"
	"github.com/Bharat1Rajput/payoutX/payout-service/internal/model"
	"github.com/Bharat1Rajput/payoutX/payout-service/internal/repository"
	"github.com/Bharat1Rajput/payoutX/payout-service/internal/transaction"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	grpcClient "github.com/Bharat1Rajput/payoutX/payout-service/internal/grpc"
)

type PayoutService struct {
	repo         repository.PayoutRepository
	ledgerClient *grpcClient.LedgerClient
	txManager    transaction.Manager
	outboxRepo   repository.OutboxRepository
}

func NewPayoutService(
	repo repository.PayoutRepository,
	ledgerClient *grpcClient.LedgerClient,
	txManager transaction.Manager,
	outboxRepo repository.OutboxRepository,
) *PayoutService {
	return &PayoutService{repo: repo, ledgerClient: ledgerClient, txManager: txManager, outboxRepo: outboxRepo}
}

func (s *PayoutService) CreatePayout(
	ctx context.Context,
	req model.CreatePayoutRequest,
) (*model.CreatePayoutResponse, error) {

	existingPayout, err := s.repo.GetByIdempotencyKey(
		ctx,
		req.IdempotencyKey,
	)

	if err == nil {
		return &model.CreatePayoutResponse{
			ID:     existingPayout.ID,
			Status: existingPayout.Status,
		}, nil
	}

	if !errors.Is(err, pgx.ErrNoRows) {
		return nil, err
	}

	payout := &model.Payout{
		ID:             uuid.NewString(),
		BeneficiaryID:  req.BeneficiaryID,
		IdempotencyKey: req.IdempotencyKey,
		Amount:         req.Amount,
		Status:         model.PayoutCreated,
		CreatedAt:      time.Now(),
	}

	err = s.ledgerClient.CreateLedgerEntries(
		ctx,
		payout.ID,
		payout.Amount,
	)

	if err != nil {
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

	err = s.txManager.WithinTransaction(
		ctx,
		func(tx pgx.Tx) error {

			err := s.repo.CreateTx(
				ctx,
				tx,
				payout,
			)

			if err != nil {
				return err
			}

			outboxEvent := &model.OutboxEvent{
				ID:        uuid.NewString(),
				Topic:     "payout.events",
				Payload:   eventBytes,
				Status:    model.OutboxPending,
				CreatedAt: time.Now(),
			}

			err = s.outboxRepo.CreateOutboxEvent(
				ctx,
				tx,
				outboxEvent,
			)

			if err != nil {
				return err
			}

			return nil
		},
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
	newStatus string,
) error {

	payout, err := s.repo.GetByID(
		ctx,
		payoutID,
	)

	if err != nil {
		return err
	}

	if !model.IsValidTransition(
		payout.Status,
		newStatus,
	) {
		return fmt.Errorf(
			"invalid state transition: %s -> %s",
			payout.Status,
			newStatus,
		)
	}

	return s.repo.UpdateStatus(
		ctx,
		payoutID,
		newStatus,
	)
}
