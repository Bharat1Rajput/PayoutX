package main

import (
	"context"

	"github.com/Bharat1Rajput/payoutX/processor-service/internal/bank"
	"github.com/Bharat1Rajput/payoutX/processor-service/internal/kafka"
	"github.com/Bharat1Rajput/payoutX/processor-service/internal/worker"
)

func main() {

	bankClient := bank.NewClient(
		"http://localhost:8081",
	)

	payoutWorker := worker.NewPayoutWorker(
		bankClient,
	)

	consumer := kafka.NewConsumer(
		"localhost:9092",
		"payout.created",
		"processor-group",
		payoutWorker,
	)

	consumer.Start(context.Background())
}
