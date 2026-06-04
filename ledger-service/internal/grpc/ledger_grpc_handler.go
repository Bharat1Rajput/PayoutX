package grpc

import (
	"context"

	pb "github.com/Bharat1Rajput/payoutX/proto/ledger"

	"github.com/Bharat1Rajput/payoutX/ledger-service/internal/model"
	"github.com/Bharat1Rajput/payoutX/ledger-service/internal/service"
)

type LedgerGRPCHandler struct {
	pb.UnimplementedLedgerServiceServer

	service *service.LedgerService
}

func NewLedgerGRPCHandler(
	service *service.LedgerService,
) *LedgerGRPCHandler {

	return &LedgerGRPCHandler{
		service: service,
	}
}

func (h *LedgerGRPCHandler) CreateLedgerEntries(
	ctx context.Context,
	req *pb.CreateLedgerEntriesRequest,
) (*pb.CreateLedgerEntriesResponse, error) {

	ledgerReq := model.CreateLedgerEntryRequest{
		TransactionID: req.TransactionId,
		Amount:        req.Amount,
	}

	err := h.service.CreateLedgerEntries(
		ctx,
		ledgerReq,
	)

	if err != nil {
		return nil, err
	}

	return &pb.CreateLedgerEntriesResponse{
		Message: "ledger entries created",
	}, nil
}