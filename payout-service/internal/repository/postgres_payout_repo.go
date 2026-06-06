package repository

import (
	"context"

	"github.com/Bharat1Rajput/payoutX/payout-service/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresPayoutRepo struct {
	db *pgxpool.Pool
}

func NewPostgresPayoutRepo(db *pgxpool.Pool) *PostgresPayoutRepo {
	return &PostgresPayoutRepo{db: db}
}

func (r *PostgresPayoutRepo) Create(ctx context.Context, payout *model.Payout) error {
	query := `
INSERT INTO	payouts 
(id,beneficiary_id,amount,status,created_at)
VALUES ($1,$2,$3,$4,$5)
`

	_, err := r.db.Exec(ctx, query, payout.ID, payout.BeneficiaryID, payout.Amount, payout.Status, payout.CreatedAt)

	return err

}

func (r *PostgresPayoutRepo) UpdateStatus(ctx context.Context, payoutID string, status string) error {
	query := `
UPDATE payouts
SET status = $1
WHERE id = $2
`
	_, err := r.db.Exec(ctx, query, status, payoutID)

	return err
}

func (r *PostgresPayoutRepo) GetByID(
	ctx context.Context,
	id string,
) (*model.Payout, error) {

	query := `
		SELECT
			id,
			beneficiary_id,
			amount,
			status,
			created_at
		FROM payouts
		WHERE id = $1
	`

	var payout model.Payout

	err := r.db.QueryRow(
		ctx,
		query,
		id,
	).Scan(
		&payout.ID,
		&payout.BeneficiaryID,
		&payout.Amount,
		&payout.Status,
		&payout.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &payout, nil
}

func (r *PostgresPayoutRepo) GetByIdempotencyKey(
	ctx context.Context,
	key string,
) (*model.Payout, error) {

	query := `
		SELECT
			id,
			beneficiary_id,
			amount,
			status,
			idempotency_key,
			created_at
		FROM payouts
		WHERE idempotency_key = $1
	`

	var payout model.Payout

	err := r.db.QueryRow(
		ctx,
		query,
		key,
	).Scan(
		&payout.ID,
		&payout.BeneficiaryID,
		&payout.Amount,
		&payout.Status,
		&payout.IdempotencyKey,
		&payout.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &payout, nil
}