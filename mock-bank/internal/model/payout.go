package model

type CreatePayoutRequest struct {
	PayoutID string `json:"payout_id"`
	Amount   int64  `json:"amount"`
}

type CreatePayoutResponse struct {
	BankReference string `json:"bank_reference"`
	Status        string `json:"status"`
}