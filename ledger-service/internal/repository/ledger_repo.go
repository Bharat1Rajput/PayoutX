package repository

import (
	"context"

	"github.com/Bharat1Rajput/payoutX/ledger-service/internal/model"
)

type LedgerRepo interface {
	CreateEntries(
		ctx context.Context,
		entries []model.LedgerEntry,
	) error
}
