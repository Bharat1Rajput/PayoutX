# PayoutX – Event-Driven Payout & Settlement Platform

PayoutX is a production-inspired fintech backend system that simulates how modern payment companies process payouts safely and reliably at scale.

Built using Go, Kafka, PostgreSQL, gRPC, and Docker, the project focuses on real fintech challenges such as ledgering, event-driven workflows, idempotency, settlement, reconciliation, webhook processing, and distributed system reliability.

## Architecture

```text
Merchant
    |
    v
Payout Service
    |
    v
Ledger Service
    |
    v
Kafka
    |
    v
Bank Processor
    |
    v
Webhook Handler
    |
    v
Settlement & Reconciliation
```

## Key Features

* Double-Entry Ledger System
* Event-Driven Architecture with Kafka
* gRPC Service-to-Service Communication
* Payout State Machine
* Webhook Processing
* Idempotency Protection
* Retry Mechanisms & DLQ
* Settlement Tracking
* Reconciliation Engine
* Transactional Outbox Pattern

## Tech Stack

* Go
* PostgreSQL
* Kafka
* gRPC
* Docker
* Gin

## Why I Built It

I wanted to understand how fintech systems handle real-world challenges such as preventing duplicate money movement, maintaining financial correctness, processing asynchronous events, and reconciling transactions across distributed services.

The project helped me learn distributed systems, financial ledgering, event-driven architecture, and backend design patterns commonly used in modern payment infrastructure.
