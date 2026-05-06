-- Add role column to users table
ALTER TABLE resocialization.users
ADD COLUMN IF NOT EXISTS role VARCHAR(20) NOT NULL DEFAULT 'user';

-- Create index on role for better query performance
CREATE INDEX IF NOT EXISTS idx_users_role ON resocialization.users(role);

-- Add comment to clarify valid role values
COMMENT ON COLUMN resocialization.users.role IS 'Role do usuário: user (padrão), admin';

-- Update existing users to have 'user' role by default
UPDATE resocialization.users SET role = 'user' WHERE role IS NULL;
