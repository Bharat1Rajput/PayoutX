package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgres() *pgxpool.Pool {
	db, err := pgxpool.New(context.Background(), "postgres://postgres:postgres@localhost:5432/payoutx")

	if err != nil {
		log.Fatal(err)
	}

	return db

}
