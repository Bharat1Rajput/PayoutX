package repository

import (
	"context"

	"github.com/Bharat1Rajput/payoutX/payout-service/internal/model"
	"github.com/jackc/pgx/v5"
)

type OutboxRepository interface {
	Create(
		ctx context.Context,
		event *model.OutboxEvent,
	) error

	CreateOutboxEvent(
		ctx context.Context,
		tx pgx.Tx,
		event *model.OutboxEvent,
	) error
	GetPending(
		ctx context.Context,
	) ([]model.OutboxEvent, error)

	MarkSent(
		ctx context.Context,
		id string,
	) error
}
