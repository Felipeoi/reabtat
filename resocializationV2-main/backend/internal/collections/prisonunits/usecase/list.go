package usecase

import (
	"context"

	"github.com/jeje-gab/resocializationV2/backend/internal/collections/prisonunits"
	"github.com/jeje-gab/resocializationV2/backend/internal/entity"
)

type ListUC struct {
	repo prisonunits.Repository
}

func NewListUC(repo prisonunits.Repository) *ListUC {
	return &ListUC{repo: repo}
}

func (uc *ListUC) Execute(ctx context.Context, ufCode string) ([]entity.PrisonUnit, error) {
	return uc.repo.List(ctx, ufCode)
}
