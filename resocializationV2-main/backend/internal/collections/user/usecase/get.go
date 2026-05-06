package usecase

import (
	domain "github.com/jeje-gab/resocializationV2/backend/internal/collections/user"
	entity "github.com/jeje-gab/resocializationV2/backend/internal/entity"
)

type GetUC struct{ repo domain.Repository }

func NewGetUC(repo domain.Repository) *GetUC { return &GetUC{repo: repo} }

func (uc *GetUC) Execute(id int64) (*entity.User, error) {
	return uc.repo.Get(id)
}
