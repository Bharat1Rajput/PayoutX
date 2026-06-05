package model

type BankWebhookRequest struct {
	PayoutID string `json:"payout_id"`
	Status   string `json:"status"`
}