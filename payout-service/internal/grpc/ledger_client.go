package grpc

import (
	"context"
	"log"
	"time"

	pb "github.com/Bharat1Rajput/payoutX/proto/ledger"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type LedgerClient struct {
	client pb.LedgerServiceClient
}

func NewLedgerClient() *LedgerClient {

	conn, err := grpc.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
	)

	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewLedgerServiceClient(conn)

	return &LedgerClient{
		client: client,
	}
}

func (l *LedgerClient) CreateLedgerEntries(
	ctx context.Context,
	transactionID string,
	amount int64,
) error {

	ctx, cancel := context.WithTimeout(
		ctx,
		5*time.Second,
	)

	defer cancel()

	_, err := l.client.CreateLedgerEntries(
		ctx,
		&pb.CreateLedgerEntriesRequest{
			TransactionId: transactionID,
			Amount:        amount,
		},
	)

	return err
}