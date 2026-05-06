package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jeje-gab/resocializationV2/backend/internal/collections/ufs"
	"github.com/jeje-gab/resocializationV2/backend/internal/entity"
)

type Repository struct {
	list *ListRepo
}

func NewRepository(pool *pgxpool.Pool) ufs.Repository {
	return &Repository{
		list: NewListRepo(pool),
	}
}

func (r *Repository) List(ctx context.Context) ([]entity.UF, error) {
	return r.list.Execute(ctx)
}
