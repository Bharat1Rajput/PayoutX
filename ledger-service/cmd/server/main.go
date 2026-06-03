package main

import (
	"context"
	"log"

	"github.com/Bharat1Rajput/ledger-service/internal/database"
	"github.com/Bharat1Rajput/ledger-service/internal/model"
	"github.com/Bharat1Rajput/ledger-service/internal/repository"
	"github.com/Bharat1Rajput/ledger-service/internal/service"
)

func main() {

	db := database.NewPostgres()

	repo := repository.NewPostgresLedgerRepo(db)

	ledgerService := service.NewLedgerService(
		repo,
	)

	req := model.CreateLedgerEntryRequest{
		TransactionID: "payout_456",
		Amount:        1000,
	}

	err := ledgerService.CreateLedgerEntries(
		context.Background(),
		req,
	)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("ledger entries created")
}
