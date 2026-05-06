CREATE SCHEMA IF NOT EXISTS resocialization;

CREATE TABLE IF NOT EXISTS resocialization.users
(
    id SERIAL PRIMARY KEY,
    name          TEXT      NOT NULL,
    email         TEXT      NOT NULL UNIQUE,
    password_hash TEXT      NOT NULL,
    created_at    TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMP NOT NULL DEFAULT NOW()
    );

CREATE INDEX IF NOT EXISTS idx_users_created_at ON resocialization.users(created_at);

CREATE TABLE IF NOT EXISTS resocialization.declarations
(
    id SERIAL PRIMARY KEY,
    student_name     TEXT        NOT NULL,
    declaration_type VARCHAR(30) NOT NULL, -- enum: matricula|conclusao|frequencia
    status           VARCHAR(20) NOT NULL DEFAULT 'RECEBIDA',
    created_at       TIMESTAMP   NOT NULL DEFAULT NOW(),
    attachment_uuid  UUID,
    s3_object_key    TEXT,                 -- path do objeto no bucket
    file_url         TEXT                  -- URL pública ou presign
    );

-- Índices para filtros
CREATE INDEX IF NOT EXISTS idx_declarations_type ON resocialization.declarations (declaration_type);
CREATE INDEX IF NOT EXISTS idx_declarations_month ON resocialization.declarations (date_trunc('month', created_at));
