// usecase/usecase.go
package usecase

import (
	"context"

	inmatespkg "github.com/jeje-gab/resocializationV2/backend/internal/collections/inmates"
	"github.com/jeje-gab/resocializationV2/backend/internal/entity"
)

type Usecase struct {
	create *CreateUC
	update *UpdateUC
	byId   *ByIdUC
	list   *ListUC
	delete *DeleteUC
}

func NewUseCase(repo inmatespkg.Repository) inmatespkg.UseCase {
	by := NewByIdUC(repo)
	return &Usecase{
		create: NewCreateUC(repo, by),
		update: NewUpdateUC(repo, by),
		byId:   by,
		list:   NewListUC(repo),
		delete: NewDeleteUC(repo),
	}
}

func (u *Usecase) Create(ctx context.Context, in entity.InmatesFlat) (error, context.Context, *entity.InmatesResp) {
	return u.create.Execute(ctx, in)
}

func (u *Usecase) ById(ctx context.Context, id int) (error, context.Context, *entity.InmatesResp) {
	return u.byId.Execute(ctx, id)
}

func (u *Usecase) Update(ctx context.Context, in entity.InmatesFlat) (error, context.Context, *entity.InmatesResp) {
	return u.update.Execute(ctx, in)
}

func (u *Usecase) List(ctx context.Context, limit, offset int) (error, context.Context, []entity.InmatesList) {
	return u.list.Execute(ctx, limit, offset)
}

func (u *Usecase) Delete(ctx context.Context, id int) (error, context.Context) {
	return u.delete.Execute(ctx, id)
}

// compile-time check
var _ inmatespkg.UseCase = (*Usecase)(nil)
