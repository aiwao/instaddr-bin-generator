CREATE TABLE IF NOT EXISTS accounts (
    id TEXT NOT NULL,
    password TEXT NOT NULL,
    amount INTEGER NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
);
