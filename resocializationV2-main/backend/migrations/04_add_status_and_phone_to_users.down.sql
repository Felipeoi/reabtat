-- Remove status and telefone columns from users table
DROP INDEX IF EXISTS resocialization.idx_users_status;

ALTER TABLE resocialization.users
DROP COLUMN IF EXISTS status,
DROP COLUMN IF EXISTS telefone;
