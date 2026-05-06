package usecase

import (
	"context"

	"github.com/jeje-gab/resocializationV2/backend/internal/collections/inmates"
)

type DeleteUC struct {
	repo inmates.Repository
}

func NewDeleteUC(repo inmates.Repository) *DeleteUC {
	return &DeleteUC{repo: repo}
}

func (u *DeleteUC) Execute(ctx context.Context, id int) (error, context.Context) {
	userID, err := getUserIDFromContext(ctx)
	if err != nil {
		return err, ctx
	}
	return u.repo.Delete(ctx, id, userID)
}
