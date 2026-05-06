package usecase

import (
	"context"

	"github.com/jeje-gab/resocializationV2/backend/internal/collections/cities"
	"github.com/jeje-gab/resocializationV2/backend/internal/entity"
)

type UseCase struct {
	list *ListUC
}

func NewUseCase(repo cities.Repository) cities.Usecase {
	return &UseCase{
		list: NewListUC(repo),
	}
}

func (u *UseCase) List(ctx context.Context, ufCode string) ([]entity.City, error) {
	return u.list.Execute(ctx, ufCode)
}
