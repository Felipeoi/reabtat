package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jeje-gab/resocializationV2/backend/internal/collections/inmates"
	"github.com/jeje-gab/resocializationV2/backend/internal/entity"
)

type Repository struct {
	create *CreateRepo
	byId   *ByIdRepo
	update *UpdateRepo
	list   *ListRepo
	delete *DeleteRepo
}

func NewRepository(pool *pgxpool.Pool) inmates.Repository {
	return &Repository{
		create: NewCreateRepo(pool),
		byId:   NewByIdRepo(pool),
		update: NewUpdateRepo(pool),
		list:   NewListRepo(pool),
		delete: NewDeleteRepo(pool),
	}
}

func (r *Repository) ById(ctx context.Context, id int, userID int64) (error, context.Context, *entity.InmatesResp) {
	return r.byId.Execute(ctx, id, userID)
}

func (r *Repository) Create(ctx context.Context, in entity.InmatesFlat) (error, context.Context, int) {
	return r.create.Execute(ctx, in)
}

func (r *Repository) Update(ctx context.Context, in entity.InmatesFlat, userID int64) (error, context.Context, int) {
	return r.update.Execute(ctx, in, userID)
}

func (r *Repository) List(ctx context.Context, limit, offset int, userID int64) (error, context.Context, []entity.InmatesList) {
	return r.list.Execute(ctx, limit, offset, userID)
}

func (r *Repository) Delete(ctx context.Context, id int, userID int64) (error, context.Context) {
	return r.delete.Execute(ctx, id, userID)
}
