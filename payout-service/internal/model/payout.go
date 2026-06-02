package model

import "time"

type Payout struct {
	ID            string    `json:"id"`
	BeneficiaryID string    `json:"beneficiary_id"`
	Amount        string    `json:"amount"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"created_at"`
}


type CreatePayoutRequest struct{
	BeneficiaryID string `json:"beneficiary_id"`
	Amount string `json:"amount"`
}

type CreatePayoutResponse struct{
 ID string `json:"id"`
 Status string `json:"status"`
}