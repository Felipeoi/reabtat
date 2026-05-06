package ufs

import (
	"context"

	"github.com/jeje-gab/resocializationV2/backend/internal/entity"
)

type Repository interface {
	List(ctx context.Context) ([]entity.UF, error)
}
