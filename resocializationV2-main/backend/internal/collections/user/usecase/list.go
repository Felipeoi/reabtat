package usecase

import (
	domain "github.com/jeje-gab/resocializationV2/backend/internal/collections/user"
	entity "github.com/jeje-gab/resocializationV2/backend/internal/entity"
)

type ListUC struct{ repo domain.Repository }

func NewListUC(repo domain.Repository) *ListUC { return &ListUC{repo: repo} }

func (uc *ListUC) Execute(limit, offset int) ([]entity.User, error) {
	return uc.repo.List(limit, offset)
}
