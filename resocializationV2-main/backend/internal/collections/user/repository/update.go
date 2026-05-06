package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UpdateRepo struct {
	pool *pgxpool.Pool
}

func NewUpdateRepo(pool *pgxpool.Pool) *UpdateRepo {
	return &UpdateRepo{
		pool: pool,
	}
}

func (r *UpdateRepo) Execute(id int64, name, email string) error {
	_, err := r.pool.Exec(context.Background(),
		`UPDATE resocialization.users SET name=$2, email=$3, updated_at=NOW() WHERE id=$1`,
		id, name, email)
	return err
}
