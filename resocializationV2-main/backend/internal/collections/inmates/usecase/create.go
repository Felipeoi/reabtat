package usecase

import (
	"context"

	"github.com/jeje-gab/resocializationV2/backend/internal/collections/inmates"
	"github.com/jeje-gab/resocializationV2/backend/internal/entity"
)

type CreateUC struct {
	repo inmates.Repository
	byId *ByIdUC
}

func NewCreateUC(repo inmates.Repository, byId *ByIdUC) *CreateUC {
	return &CreateUC{repo: repo, byId: byId}
}

func (u *CreateUC) Execute(ctx context.Context, in entity.InmatesFlat) (error, context.Context, *entity.InmatesResp) {
	// Extrai o user_id do contexto (setado pelo handler a partir do JWT)
	userID, err := getUserIDFromContext(ctx)
	if err != nil {
		return err, ctx, nil
	}

	// Injeta o user_id no objeto
	in.UserID = userID

	// cria e pega o ID
	err, ctx, createdID := u.repo.Create(ctx, in)
	if err != nil {
		return err, ctx, nil
	}
	// carrega o registro completo
	err, ctx, resp := u.byId.Execute(ctx, createdID)
	return err, ctx, resp
}
