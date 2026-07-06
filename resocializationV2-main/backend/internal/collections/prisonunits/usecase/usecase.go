package usecase

import (
	"context"

	"github.com/jeje-gab/resocializationV2/backend/internal/collections/prisonunits"
	"github.com/jeje-gab/resocializationV2/backend/internal/entity"
)

type UseCase struct {
	list *ListUC
}

func NewUseCase(repo prisonunits.Repository) prisonunits.Usecase {
	return &UseCase{list: NewListUC(repo)}
}

func (u *UseCase) List(ctx context.Context, ufCode string) ([]entity.PrisonUnit, error) {
	return u.list.Execute(ctx, ufCode)
}
