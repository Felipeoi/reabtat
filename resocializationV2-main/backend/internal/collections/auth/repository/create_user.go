package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CreateUserRepo struct {
	pool *pgxpool.Pool
}

func NewCreateUserRepo(pool *pgxpool.Pool) *CreateUserRepo {
	return &CreateUserRepo{
		pool: pool,
	}
}

func (r *CreateUserRepo) Execute(name, email, passwordHash, telefone string) (int64, error) {
	var id int64
	err := r.pool.QueryRow(context.Background(),
		`INSERT INTO resocialization.users (name, email, password_hash, telefone, status, role) VALUES ($1,$2,$3,$4,DEFAULT,'user') RETURNING id`,
		name, email, passwordHash, telefone).Scan(&id)
	return id, err
}
