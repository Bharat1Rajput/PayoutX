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
