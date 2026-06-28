package main

import (
	"github.com/gin-gonic/gin"

	"github.com/Bharat1Rajput/payoutX/mock-bank/internal/handler"
	"github.com/Bharat1Rajput/payoutX/mock-bank/internal/service"
)

func main() {

	payoutService := service.NewPayoutService()

	payoutHandler := handler.NewPayoutHandler(
		payoutService,
	)

	router := gin.Default()

	router.POST(
		"/payouts",
		payoutHandler.CreatePayout,
	)

	router.GET(
		"/payouts/:id",
		payoutHandler.GetPayout,
	)
	router.Run(":8081")
}
