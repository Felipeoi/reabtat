package usecase

import (
	"context"

	"github.com/jeje-gab/resocializationV2/backend/internal/collections/ufs"
	"github.com/jeje-gab/resocializationV2/backend/internal/entity"
)

type UseCase struct {
	list *ListUC
}

func NewUseCase(repo ufs.Repository) ufs.Usecase {
	return &UseCase{
		list: NewListUC(repo),
	}
}

func (u *UseCase) List(ctx context.Context) ([]entity.UF, error) {
	return u.list.Execute(ctx)
}
