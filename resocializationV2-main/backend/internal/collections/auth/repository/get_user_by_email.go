package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jeje-gab/resocializationV2/backend/internal/entity"
)

type GetUserByEmailRepo struct {
	pool *pgxpool.Pool
}

func NewGetUserByEmailRepo(pool *pgxpool.Pool) *GetUserByEmailRepo {
	return &GetUserByEmailRepo{pool: pool}
}

func (r *GetUserByEmailRepo) Execute(email string) (*entity.UserAuth, error) {
	row := r.pool.QueryRow(context.Background(),
		`SELECT id, name, email, password_hash, status, role
		 FROM resocialization.users
		 WHERE LOWER(TRIM(email)) = LOWER(TRIM($1::text))
		 LIMIT 1`, email)
	var u entity.UserAuth
	if err := row.Scan(&u.ID, &u.Name, &u.Email, &u.PasswordHash, &u.Status, &u.Role); err != nil {
		return nil, err
	}
	return &u, nil
}
