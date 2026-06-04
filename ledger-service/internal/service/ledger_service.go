package service

import (
	"context"
	"time"

	"github.com/Bharat1Rajput/payoutX/ledger-service/internal/model"
	"github.com/Bharat1Rajput/payoutX/ledger-service/internal/repository"
	"github.com/google/uuid"
)

type LedgerService struct {
	repo repository.LedgerRepo
}

func NewLedgerService(repo repository.LedgerRepo) *LedgerService {
	return &LedgerService{repo: repo}
}

func (s *LedgerService) CreateLedgerEntries(ctx context.Context, req model.CreateLedgerEntryRequest) error {

	entries := []model.LedgerEntry{
		{
			ID:            uuid.NewString(),
			TransactionID: req.TransactionID,
			Account:       model.AccountMerchantWallet,
			EntryType:     model.EntryTypeDebit,
			Amount:        req.Amount,
			CreatedAt:     time.Now(),
		},
		{
			ID:            uuid.NewString(),
			TransactionID: req.TransactionID,
			Account:       model.AccountPendingPayout,
			EntryType:     model.EntryTypeCredit,
			Amount:        req.Amount,
			CreatedAt:     time.Now(),
		},
	}

	return s.repo.CreateEntries(ctx, entries)

}
