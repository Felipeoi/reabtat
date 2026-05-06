package usecase

import (
	"context"

	"github.com/jeje-gab/resocializationV2/backend/internal/collections/inmates"
	"github.com/jeje-gab/resocializationV2/backend/internal/entity"
)

type UpdateUC struct {
	repo inmates.Repository
	byId *ByIdUC
}

func NewUpdateUC(repo inmates.Repository, byId *ByIdUC) *UpdateUC {
	return &UpdateUC{repo: repo, byId: byId}
}

func (u *UpdateUC) Execute(ctx context.Context, in entity.InmatesFlat) (error, context.Context, *entity.InmatesResp) {
	userID, err := getUserIDFromContext(ctx)
	if err != nil {
		return err, ctx, nil
	}
	err, ctx, updatedID := u.repo.Update(ctx, in, userID)
	if err != nil {
		return err, ctx, nil
	}
	err, ctx, resp := u.byId.Execute(ctx, updatedID)
	return err, ctx, resp
}
