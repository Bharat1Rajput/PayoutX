package main

import (
	"log"
	"net"

	pb "github.com/Bharat1Rajput/payoutX/proto/ledger"

	grpcHandler "github.com/Bharat1Rajput/payoutX/ledger-service/internal/grpc"

	"github.com/Bharat1Rajput/payoutX/ledger-service/internal/database"
	"github.com/Bharat1Rajput/payoutX/ledger-service/internal/repository"
	"github.com/Bharat1Rajput/payoutX/ledger-service/internal/service"

	"google.golang.org/grpc"
)

func main() {

	db := database.NewPostgres()

	repo := repository.NewPostgresLedgerRepo(
		db,
	)

	ledgerService := service.NewLedgerService(
		repo,
	)

	handler := grpcHandler.NewLedgerGRPCHandler(
		ledgerService,
	)

	lis, err := net.Listen(
		"tcp",
		":50051",
	)

	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterLedgerServiceServer(
		grpcServer,
		handler,
	)

	log.Println("ledger grpc server running on :50051")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
