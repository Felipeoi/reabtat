package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jeje-gab/resocializationV2/backend/internal/entity"
)

type ListRepo struct {
	pool *pgxpool.Pool
}

func NewListRepo(pool *pgxpool.Pool) *ListRepo {
	return &ListRepo{
		pool: pool,
	}
}

func (r *ListRepo) Execute(limit, offset int) ([]entity.User, error) {
	rows, err := r.pool.Query(context.Background(),
		`SELECT id, name, email, status, role, telefone, created_at, updated_at
		   FROM resocialization.users ORDER BY created_at DESC
		   LIMIT $1 OFFSET $2`, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []entity.User
	for rows.Next() {
		var u entity.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Status, &u.Role, &u.Telefone, &u.CreatedAt, &u.UpdatedAt); err != nil {
			return nil, err
		}
		out = append(out, u)
	}
	return out, rows.Err()
}
