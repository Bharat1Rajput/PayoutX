package model

import "time"

type OutboxEvent struct {
	ID        string
	Topic     string
	Payload   []byte
	Status    string
	CreatedAt time.Time
}


const (
	OutboxPending = "PENDING"
	OutboxSent    = "SENT"
)