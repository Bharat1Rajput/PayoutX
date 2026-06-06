CREATE TABLE IF NOT EXISTS payouts (
    id UUID PRIMARY KEY,
    beneficiary_id TEXT NOT NULL,
    idempotency_key TEXT UNIQUE,
    amount BIGINT NOT NULL,
    status TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL
);