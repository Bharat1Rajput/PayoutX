package main

import (
	"github.com/gin-gonic/gin"

	"github.com/Bharat1Rajput/payout-service/internal/database"
	"github.com/Bharat1Rajput/payout-service/internal/handler"
	"github.com/Bharat1Rajput/payout-service/internal/repository"
	"github.com/Bharat1Rajput/payout-service/internal/service"
)

func main() {

	db := database.NewPostgres()

	repo := repository.NewPostgresPayoutRepo(db)

	payoutService := service.NewPayoutService(repo)

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
