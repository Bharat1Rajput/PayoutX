package transaction

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresManager struct {
	db *pgxpool.Pool
}

func NewPostgresManager(
	db *pgxpool.Pool,
) *PostgresManager {

	return &PostgresManager{
		db: db,
	}
}

func (m *PostgresManager) WithinTransaction(
	ctx context.Context,
	fn func(tx pgx.Tx) error,
) error {

	tx, err := m.db.Begin(ctx)
	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)

	err = fn(tx)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}
