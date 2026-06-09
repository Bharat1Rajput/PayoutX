package repository

import (
	"context"

	"github.com/Bharat1Rajput/payoutX/payout-service/internal/model"
	"github.com/jackc/pgx/v5"
)

type PayoutRepository interface {
	Create(
		ctx context.Context, payout *model.Payout,
	) error
	UpdateStatus(
		ctx context.Context, payoutID string, status string,
	) error

	GetByID(
		ctx context.Context,
		id string,
	) (*model.Payout, error)

	GetByIdempotencyKey(
		ctx context.Context,
		key string,
	) (*model.Payout, error)

	CreateTx(
		ctx context.Context,
		tx pgx.Tx,
		payout *model.Payout,
	) error
}
