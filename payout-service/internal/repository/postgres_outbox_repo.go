package repository

import (
	"context"

	"github.com/Bharat1Rajput/payoutX/payout-service/internal/model"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresOutboxRepo struct {
	db *pgxpool.Pool
}

func NewPostgresOutboxRepo(
	db *pgxpool.Pool,
) *PostgresOutboxRepo {

	return &PostgresOutboxRepo{
		db: db,
	}
}

func (r *PostgresOutboxRepo) Create(
	ctx context.Context,
	event *model.OutboxEvent,
) error {

	query := `
		INSERT INTO outbox_events (
			id,
			topic,
			payload,
			status,
			created_at
		)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err := r.db.Exec(
		ctx,
		query,
		event.ID,
		event.Topic,
		event.Payload,
		event.Status,
		event.CreatedAt,
	)

	return err
}

func (r *PostgresOutboxRepo) CreateOutboxEvent(
	ctx context.Context,
	tx pgx.Tx,
	event *model.OutboxEvent,
) error {
 
	query := `
		INSERT INTO outbox_events
		(
			id,
			topic,
			payload,
			status,
			created_at
		)
		VALUES ($1,$2,$3,$4,$5)
	`

	_, err := tx.Exec(
		ctx,
		query,
		event.ID,
		event.Topic,
		event.Payload,
		event.Status,
		event.CreatedAt,
	)

	return err
}

func (r *PostgresOutboxRepo) GetPending(
	ctx context.Context,
) ([]model.OutboxEvent, error) {

	query := `
		SELECT
			id,
			topic,
			payload,
			status,
			created_at
		FROM outbox_events
		WHERE status = $1
		ORDER BY created_at ASC
	`

	rows, err := r.db.Query(
		ctx,
		query,
		model.OutboxPending,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []model.OutboxEvent

	for rows.Next() {

		var event model.OutboxEvent

		err := rows.Scan(
			&event.ID,
			&event.Topic,
			&event.Payload,
			&event.Status,
			&event.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		events = append(
			events,
			event,
		)
	}

	return events, nil
}

func (r *PostgresOutboxRepo) MarkSent(
	ctx context.Context,
	id string,
) error {

	query := `
		UPDATE outbox_events
		SET status = $1
		WHERE id = $2
	`

	_, err := r.db.Exec(
		ctx,
		query,
		model.OutboxSent,
		id,
	)

	return err
}