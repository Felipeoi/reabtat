package usecase

import (
	"context"

	"github.com/jeje-gab/resocializationV2/backend/internal/collections/ufs"
	"github.com/jeje-gab/resocializationV2/backend/internal/entity"
)

type ListUC struct {
	repo ufs.Repository
}

func NewListUC(repo ufs.Repository) *ListUC {
	return &ListUC{repo: repo}
}

func (uc *ListUC) Execute(ctx context.Context) ([]entity.UF, error) {
	return uc.repo.List(ctx)
}
