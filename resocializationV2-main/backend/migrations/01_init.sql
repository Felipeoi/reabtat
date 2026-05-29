CREATE SCHEMA IF NOT EXISTS resocialization;

CREATE TABLE IF NOT EXISTS resocialization.users
(
    id            SERIAL PRIMARY KEY,
    name          TEXT      NOT NULL,
    email         TEXT      NOT NULL UNIQUE,
    password_hash TEXT      NOT NULL,
    created_at    TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_users_created_at ON resocialization.users (created_at);
