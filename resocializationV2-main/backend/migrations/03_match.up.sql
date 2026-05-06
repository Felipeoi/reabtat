-- Tabela de Matches (para futuras funcionalidades de aceitar/rejeitar matches)
-- Por enquanto, os matches são calculados em tempo real através de queries
-- Esta tabela pode ser usada no futuro para armazenar matches aceitos/propostos

CREATE TABLE IF NOT EXISTS resocialization.matches
(
    id              SERIAL PRIMARY KEY,
    inmate_a_id     INTEGER   NOT NULL REFERENCES resocialization.inmates (id) ON DELETE CASCADE,
    inmate_b_id     INTEGER   NOT NULL REFERENCES resocialization.inmates (id) ON DELETE CASCADE,
    status          VARCHAR(20) NOT NULL DEFAULT 'PENDING' CHECK (status IN ('PENDING', 'ACCEPTED', 'REJECTED')),
    proposed_by     INTEGER   NOT NULL REFERENCES resocialization.users (id) ON DELETE CASCADE,
    proposed_at     TIMESTAMP NOT NULL DEFAULT NOW(),
    responded_at    TIMESTAMP,

    -- Garante que não haja duplicatas (match A-B é o mesmo que B-A)
    CONSTRAINT matches_unique CHECK (inmate_a_id < inmate_b_id),
    CONSTRAINT matches_unique_pair UNIQUE (inmate_a_id, inmate_b_id)
);

-- Índices para melhorar performance
CREATE INDEX IF NOT EXISTS idx_matches_inmate_a ON resocialization.matches(inmate_a_id);
CREATE INDEX IF NOT EXISTS idx_matches_inmate_b ON resocialization.matches(inmate_b_id);
CREATE INDEX IF NOT EXISTS idx_matches_status ON resocialization.matches(status);
CREATE INDEX IF NOT EXISTS idx_matches_proposed_by ON resocialization.matches(proposed_by);

COMMENT ON TABLE resocialization.matches IS 'Armazena propostas de match entre inmates para troca de localização';
COMMENT ON COLUMN resocialization.matches.inmate_a_id IS 'ID do primeiro inmate (sempre menor que inmate_b_id)';
COMMENT ON COLUMN resocialization.matches.inmate_b_id IS 'ID do segundo inmate (sempre maior que inmate_a_id)';
COMMENT ON COLUMN resocialization.matches.status IS 'Status do match: PENDING, ACCEPTED, REJECTED';
COMMENT ON COLUMN resocialization.matches.proposed_by IS 'ID do usuário que propôs o match';
