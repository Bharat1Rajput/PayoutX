package kafka

type PayoutCreatedEvent struct {
	PayoutID     string `json:"payout_id"`
	BeneficiaryID string `json:"beneficiary_id"`
	Amount       int64  `json:"amount"`
	Status       string `json:"status"`
}