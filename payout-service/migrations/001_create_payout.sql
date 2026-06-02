CREATE TABLE IF NOT EXISTS payouts (
    id UUID PRIMARY KEY,
    beneficiary_id TEXT NOT NULL,
    amount BIGINT NOT NULL,
    status TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL
);