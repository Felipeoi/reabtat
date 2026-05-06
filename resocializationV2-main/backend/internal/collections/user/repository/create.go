package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CreateRepo struct {
	pool *pgxpool.Pool
}

func NewCreateRepo(pool *pgxpool.Pool) *CreateRepo {
	return &CreateRepo{
		pool: pool,
	}
}

func (r *CreateRepo) Execute(name, email, passwordHash string) (int64, error) {
	var id int64
	err := r.pool.QueryRow(context.Background(),
		`INSERT INTO resocialization.users (name, email, password_hash) VALUES ($1,$2,$3) RETURNING id`,
		name, email, passwordHash).Scan(&id)
	return id, err
}
