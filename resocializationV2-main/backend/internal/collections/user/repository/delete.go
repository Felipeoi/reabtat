package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DeleteRepo struct {
	pool *pgxpool.Pool
}

func NewDeleteRepo(pool *pgxpool.Pool) *DeleteRepo {
	return &DeleteRepo{
		pool: pool,
	}
}

func (r *DeleteRepo) Execute(id int64) error {
	_, err := r.pool.Exec(context.Background(), `DELETE FROM resocialization.users WHERE id=$1`, id)
	return err
}
