package usecase

import (
	"context"

	"github.com/jeje-gab/resocializationV2/backend/internal/collections/cities"
	"github.com/jeje-gab/resocializationV2/backend/internal/entity"
)

type ListUC struct {
	repo cities.Repository
}

func NewListUC(repo cities.Repository) *ListUC {
	return &ListUC{
		repo: repo,
	}
}

func (u *ListUC) Execute(ctx context.Context, ufCode string) ([]entity.City, error) {
	return u.repo.List(ctx, ufCode)
}
