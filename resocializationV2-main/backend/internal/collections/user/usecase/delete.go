package usecase

import (
	domain "github.com/jeje-gab/resocializationV2/backend/internal/collections/user"
)

type DeleteUC struct{ repo domain.Repository }

func NewDeleteUC(repo domain.Repository) *DeleteUC { return &DeleteUC{repo: repo} }

func (uc *DeleteUC) Execute(id int64) error {
	return uc.repo.Delete(id)
}
