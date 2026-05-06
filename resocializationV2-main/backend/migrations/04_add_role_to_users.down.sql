-- Remove role column from users table
DROP INDEX IF EXISTS resocialization.idx_users_role;

ALTER TABLE resocialization.users
DROP COLUMN IF EXISTS role;
