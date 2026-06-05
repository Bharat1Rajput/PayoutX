package main

import (
	"github.com/gin-gonic/gin"

	"github.com/Bharat1Rajput/payoutX/payout-service/internal/database"
	grpcClient "github.com/Bharat1Rajput/payoutX/payout-service/internal/grpc"
	"github.com/Bharat1Rajput/payoutX/payout-service/internal/handler"
	"github.com/Bharat1Rajput/payoutX/payout-service/internal/kafka"
	"github.com/Bharat1Rajput/payoutX/payout-service/internal/repository"
	"github.com/Bharat1Rajput/payoutX/payout-service/internal/service"
)

func main() {

	db := database.NewPostgres()

	repo := repository.NewPostgresPayoutRepo(db)

	ledgerClient := grpcClient.NewLedgerClient()

	kafkaProducer := kafka.NewProducer(
		"localhost:9092",
		"payout.created",
	)

	payoutService := service.NewPayoutService(
		repo,
		ledgerClient,
		kafkaProducer,
	)

	payoutHandler := handler.NewPayoutHandler(
		payoutService,
	)

	router := gin.Default()

	router.POST(
		"/payouts",
		payoutHandler.CreatePayout,
	)

	router.Run(":8080")
}
