package repository

import (
	"context"

	"github.com/Bharat1Rajput/payoutX/ledger-service/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresLedgerRepo struct {
	db *pgxpool.Pool
}

func NewPostgresLedgerRepo(db *pgxpool.Pool) *PostgresLedgerRepo {
	return &PostgresLedgerRepo{db: db}
}

func (r *PostgresLedgerRepo) CreateEntries(ctx context.Context, entries []model.LedgerEntry) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)

	query := `
	INSERT INTO ledger_entries
	(id, transaction_id, account, entry_type, amount, created_at)
	VALUES ($1, $2, $3, $4, $5, $6)
	`

	for _, entry := range entries {

		_, err := tx.Exec(
			ctx,
			query,
			entry.ID,
			entry.TransactionID,
			entry.Account,
			entry.EntryType,
			entry.Amount,
			entry.CreatedAt,
		)

		if err != nil {
			return err
		}
	}

	return tx.Commit(ctx)
}
