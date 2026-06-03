package model

import "time"

type LedgerEntry struct {
	ID            string    `json:"id"`
	TransactionID string    `json:"transaction_id"`
	Account       string    `json:"account"`
	EntryType     string    `json:"entry_type"`
	Amount        int64     `json:"amount"`
	CreatedAt     time.Time `json:"created_at"`
}

type CreateLedgerEntryRequest struct {
	TransactionID string `json:"transaction_id"`
	Amount        int64  `json:"amount"`
}