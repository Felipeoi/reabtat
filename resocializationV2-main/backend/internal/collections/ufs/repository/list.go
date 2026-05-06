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

func (r *ListRepo) Execute(ctx context.Context) ([]entity.UF, error) {
	const q = `SELECT id, code, name FROM public.ufs ORDER BY name`

	rows, err := r.pool.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]entity.UF, 0)
	for rows.Next() {
		var uf entity.UF
		if err := rows.Scan(&uf.ID, &uf.Code, &uf.Name); err != nil {
			return nil, err
		}
		result = append(result, uf)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
