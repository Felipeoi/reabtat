package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jeje-gab/resocializationV2/backend/internal/entity"
)

type GetRepo struct {
	pool *pgxpool.Pool
}

func NewGetRepo(pool *pgxpool.Pool) *GetRepo {
	return &GetRepo{
		pool: pool,
	}
}

func (r *GetRepo) Execute(id int64) (*entity.User, error) {
	row := r.pool.QueryRow(context.Background(),
		`SELECT id, name, email, status, role, telefone, created_at, updated_at FROM resocialization.users WHERE id=$1`, id)
	var u entity.User
	if err := row.Scan(&u.ID, &u.Name, &u.Email, &u.Status, &u.Role, &u.Telefone, &u.CreatedAt, &u.UpdatedAt); err != nil {
		return nil, err
	}
	return &u, nil
}
