CREATE TABLE IF NOT EXISTS payouts (
    id UUID PRIMARY KEY,
    beneficiary_id TEXT NOT NULL,
    idempotency_key TEXT UNIQUE,
    bank_reference TEXT,
    amount BIGINT NOT NULL,
    status TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL
);



CREATE TABLE outbox_events (
    id UUID PRIMARY KEY,
    topic TEXT NOT NULL,
    payload JSONB NOT NULL,
    status TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL
);




-- docker exec -it payoutx-postgres psql -U admin -d payoutx

