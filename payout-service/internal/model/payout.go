package model

import "time"

const (
	PayoutCreated    = "CREATED"
	PayoutProcessing = "PROCESSING"
	PayoutSuccess    = "SUCCESS"
	PayoutFailed     = "FAILED"
)

type Payout struct {
	ID             string    `json:"id"`
	BeneficiaryID  string    `json:"beneficiary_id"`
	IdempotencyKey string    `json:"idempotency_key,omitempty"`
	BankReference  string    `json:"bank_reference"`
	Amount         int64     `json:"amount"`
	Status         string    `json:"status"`
	CreatedAt      time.Time `json:"created_at"`
}

type CreatePayoutRequest struct {
	BeneficiaryID  string `json:"beneficiary_id"`
	IdempotencyKey string `json:"idempotency_key,omitempty"`
	Amount         int64  `json:"amount"`
}

type CreatePayoutResponse struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

type UpdateStatusRequest struct {
	Status string `json:"status"`
}

type UpdateBankReferenceRequest struct {
	BankReference string `json:"bank_reference"`
}
