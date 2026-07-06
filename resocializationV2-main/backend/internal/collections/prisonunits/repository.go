package prisonunits

import (
	"context"

	"github.com/jeje-gab/resocializationV2/backend/internal/entity"
)

type Repository interface {
	List(ctx context.Context, ufCode string) ([]entity.PrisonUnit, error)
}

type Usecase interface {
	List(ctx context.Context, ufCode string) ([]entity.PrisonUnit, error)
}
