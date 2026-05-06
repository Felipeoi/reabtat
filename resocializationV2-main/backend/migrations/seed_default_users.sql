-- Usuários de demonstração (desenvolvimento). Senha para ambos: Advogado@123
-- Execute após as migrations de users (role, status, telefone).
-- Em conflito de e-mail, atualiza senha e status para você conseguir entrar de novo
-- (ex.: já tinha cadastrado esse e-mail com outra senha).

INSERT INTO resocialization.users (name, email, password_hash, telefone, status, role)
VALUES (
    'Administrador',
    'admin@local.dev',
    '$2a$10$9hpgb6Rhp7hbq2oYFrCb3OZVLd4uHt1MKaWsoV18oTLoqGz1DaXyi',
    '',
    'ativo',
    'admin'
)
ON CONFLICT (email) DO UPDATE SET
    password_hash = EXCLUDED.password_hash,
    name            = EXCLUDED.name,
    telefone        = EXCLUDED.telefone,
    status          = EXCLUDED.status,
    role            = EXCLUDED.role;

INSERT INTO resocialization.users (name, email, password_hash, telefone, status, role)
VALUES (
    'Advogado demonstração',
    'advogado@local.dev',
    '$2a$10$9hpgb6Rhp7hbq2oYFrCb3OZVLd4uHt1MKaWsoV18oTLoqGz1DaXyi',
    '',
    'ativo',
    'user'
)
ON CONFLICT (email) DO UPDATE SET
    password_hash = EXCLUDED.password_hash,
    name            = EXCLUDED.name,
    telefone        = EXCLUDED.telefone,
    status          = EXCLUDED.status,
    role            = EXCLUDED.role;
