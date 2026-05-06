package usecase

import (
	"context"

	"github.com/jeje-gab/resocializationV2/backend/internal/collections/inmates"
	"github.com/jeje-gab/resocializationV2/backend/internal/entity"
)

type ListUC struct {
	repo inmates.Repository
}

func NewListUC(repo inmates.Repository) *ListUC {
	return &ListUC{repo: repo}
}

func (u *ListUC) Execute(ctx context.Context, limit, offset int) (error, context.Context, []entity.InmatesList) {
	userID, err := getUserIDFromContext(ctx)
	if err != nil {
		return err, ctx, nil
	}
	return u.repo.List(ctx, limit, offset, userID)
}
