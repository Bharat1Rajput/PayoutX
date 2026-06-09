package model

const (
	PayoutProcessing = "PROCESSING"
	PayoutFailed     = "FAILED"
)

type PayoutCreatedEvent struct {
	PayoutID      string `json:"payout_id"`
	BeneficiaryID string `json:"beneficiary_id"`
	Amount        int64  `json:"amount"`
	Status        string `json:"status"`
}

type PayoutFailedEvent struct {
	PayoutID string `json:"payout_id"`
	Amount   int64  `json:"amount"`
	Reason   string `json:"reason"`
}
