package usecase

import (
	"context"

	"github.com/jeje-gab/resocializationV2/backend/internal/collections/inmates"
	"github.com/jeje-gab/resocializationV2/backend/internal/entity"
)

type ByIdUC struct {
	repo inmates.Repository
}

func NewByIdUC(repo inmates.Repository) *ByIdUC {
	return &ByIdUC{repo: repo}
}

func (u *ByIdUC) Execute(ctx context.Context, id int) (error, context.Context, *entity.InmatesResp) {
	userID, err := getUserIDFromContext(ctx)
	if err != nil {
		return err, ctx, nil
	}
	return u.repo.ById(ctx, id, userID)
}
