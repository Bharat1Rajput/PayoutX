package main

import (
	"context"

	"github.com/gin-gonic/gin"

	"github.com/Bharat1Rajput/payoutX/payout-service/internal/database"
	grpcClient "github.com/Bharat1Rajput/payoutX/payout-service/internal/grpc"
	"github.com/Bharat1Rajput/payoutX/payout-service/internal/handler"
	"github.com/Bharat1Rajput/payoutX/payout-service/internal/kafka"
	"github.com/Bharat1Rajput/payoutX/payout-service/internal/outbox"
	"github.com/Bharat1Rajput/payoutX/payout-service/internal/repository"
	"github.com/Bharat1Rajput/payoutX/payout-service/internal/service"
	"github.com/Bharat1Rajput/payoutX/payout-service/internal/transaction"
)

func main() {

	db := database.NewPostgres()

	repo := repository.NewPostgresPayoutRepo(db)
	outboxRepo := repository.NewPostgresOutboxRepo(db)
	txManager := transaction.NewPostgresManager(db)

	outboxProducer := kafka.NewProducer("localhost:9092", "payout.created")
	publisher := outbox.NewPublisher(outboxRepo, outboxProducer)

	ledgerClient := grpcClient.NewLedgerClient()

	payoutService := service.NewPayoutService(
		repo,
		ledgerClient,
		txManager,
		outboxRepo,
	)

	payoutHandler := handler.NewPayoutHandler(
		payoutService,
	)

	router := gin.Default()

	router.POST(
		"/payouts",
		payoutHandler.CreatePayout,
	)
	router.PATCH(
		"/payouts/:id/status",
		payoutHandler.UpdatePayoutStatus,
	)

	router.POST(
		"/webhooks/bank",
		payoutHandler.HandleBankWebhook,
	)
	router.PATCH(
		"/payouts/:id/bank-reference",
		payoutHandler.UpdateBankReference,
	)	

	bankClient := bank.NewClient(
	"http://localhost:8081",
)

reconciliationService := reconciliation.NewService(
	payoutRepo,
	bankClient,
)

	ctx := context.Background()

	go publisher.Start(ctx)
go reconciliationService.Start(ctx)
	router.Run(":8080")
}
