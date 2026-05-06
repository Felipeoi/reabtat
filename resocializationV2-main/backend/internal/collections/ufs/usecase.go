package ufs

import (
	"context"

	"github.com/jeje-gab/resocializationV2/backend/internal/entity"
)

type Usecase interface {
	List(ctx context.Context) ([]entity.UF, error)
}
