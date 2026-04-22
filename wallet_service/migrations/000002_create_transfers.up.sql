CREATE TABLE transfers (
    id         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    from_id    UUID NOT NULL REFERENCES accounts(id),
    to_id      UUID NOT NULL REFERENCES accounts(id),
    amount     NUMERIC(20, 8) NOT NULL CHECK (amount > 0),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE INDEX idx_transfers_from_id ON transfers(from_id);
CREATE INDEX idx_transfers_to_id   ON transfers(to_id);
