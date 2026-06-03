package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgres() *pgxpool.Pool {

	connString := "postgres://admin:admin@localhost:5433/payoutx?sslmode=disable"

	db, err := pgxpool.New(
		context.Background(),
		connString,
	)

	if err != nil {
		log.Fatalf("unable to create connection pool: %v", err)
	}

	err = db.Ping(context.Background())
	if err != nil {
		log.Fatalf("unable to ping database: %v", err)
	}

	log.Println("postgres connected")

	return db
}
