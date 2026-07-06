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
	return &ListRepo{pool: pool}
}

func (r *ListRepo) Execute(ctx context.Context, ufCode string) ([]entity.PrisonUnit, error) {
	var q string
	var args []interface{}

	if ufCode != "" {
		q = `SELECT id, name, uf_code FROM public.prison_units WHERE uf_code = $1 ORDER BY name`
		args = []interface{}{ufCode}
	} else {
		q = `SELECT id, name, uf_code FROM public.prison_units ORDER BY uf_code, name`
		args = []interface{}{}
	}

	rows, err := r.pool.Query(ctx, q, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]entity.PrisonUnit, 0)
	for rows.Next() {
		var unit entity.PrisonUnit
		if err := rows.Scan(&unit.ID, &unit.Name, &unit.UFCode); err != nil {
			return nil, err
		}
		result = append(result, unit)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
