-- Add status and telefone columns to users table
ALTER TABLE resocialization.users
ADD COLUMN IF NOT EXISTS status VARCHAR(20) NOT NULL DEFAULT 'ativo',
ADD COLUMN IF NOT EXISTS telefone VARCHAR(20);

-- Create index on status for better query performance
CREATE INDEX IF NOT EXISTS idx_users_status ON resocialization.users(status);

-- Add comment to clarify valid status values
COMMENT ON COLUMN resocialization.users.status IS 'Status do usuário: ativo, inativo, bloqueado';
COMMENT ON COLUMN resocialization.users.telefone IS 'Número de telefone com código do país (ex: +5511999999999)';
