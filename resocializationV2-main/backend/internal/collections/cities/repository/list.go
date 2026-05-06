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

func (r *ListRepo) Execute(ctx context.Context, ufCode string) ([]entity.City, error) {
	var q string
	var args []interface{}

	if ufCode != "" {
		q = `SELECT id, ibge_code, name, uf_code FROM public.cities WHERE uf_code = $1 ORDER BY name`
		args = []interface{}{ufCode}
	} else {
		q = `SELECT id, ibge_code, name, uf_code FROM public.cities ORDER BY name`
		args = []interface{}{}
	}

	rows, err := r.pool.Query(ctx, q, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []entity.City
	for rows.Next() {
		var city entity.City
		if err := rows.Scan(&city.ID, &city.IBGECode, &city.Name, &city.UFCode); err != nil {
			return nil, err
		}
		result = append(result, city)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
