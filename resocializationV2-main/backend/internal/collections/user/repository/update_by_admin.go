package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UpdateByAdminRepo struct {
	pool *pgxpool.Pool
}

func NewUpdateByAdminRepo(pool *pgxpool.Pool) *UpdateByAdminRepo {
	return &UpdateByAdminRepo{
		pool: pool,
	}
}

// Execute atualiza usuário (admin pode alterar status, email e senha)
func (r *UpdateByAdminRepo) Execute(id int64, name, email, status string, passwordHash *string) error {
	// Se tem password, atualiza também
	if passwordHash != nil && *passwordHash != "" {
		_, err := r.pool.Exec(context.Background(),
			`UPDATE resocialization.users
			 SET name=$2, email=$3, status=$4, password_hash=$5, updated_at=NOW()
			 WHERE id=$1`,
			id, name, email, status, *passwordHash)
		return err
	}

	// Senão, atualiza sem password
	_, err := r.pool.Exec(context.Background(),
		`UPDATE resocialization.users
		 SET name=$2, email=$3, status=$4, updated_at=NOW()
		 WHERE id=$1`,
		id, name, email, status)
	return err
}
