CREATE TABLE IF NOT EXISTS ledger_entries (

    id UUID PRIMARY KEY,

    transaction_id TEXT NOT NULL,

    account TEXT NOT NULL,

    entry_type TEXT NOT NULL,

    amount BIGINT NOT NULL,

    created_at TIMESTAMP NOT NULL 
);