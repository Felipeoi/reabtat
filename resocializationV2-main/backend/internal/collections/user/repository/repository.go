package repository

import (
	"github.com/jeje-gab/resocializationV2/backend/internal/entity"

	"github.com/jackc/pgx/v5/pgxpool"
	domain "github.com/jeje-gab/resocializationV2/backend/internal/collections/user"
)

type Repository struct {
	create        *CreateRepo
	get           *GetRepo
	list          *ListRepo
	update        *UpdateRepo
	updateByAdmin *UpdateByAdminRepo
	delete        *DeleteRepo
}

func NewRepository(pool *pgxpool.Pool) domain.Repository {
	return &Repository{
		create:        NewCreateRepo(pool),
		get:           NewGetRepo(pool),
		list:          NewListRepo(pool),
		update:        NewUpdateRepo(pool),
		updateByAdmin: NewUpdateByAdminRepo(pool),
		delete:        NewDeleteRepo(pool),
	}
}

func (r *Repository) Create(name, email, passwordHash string) (int64, error) {
	return r.create.Execute(name, email, passwordHash)
}

func (r *Repository) Get(id int64) (*entity.User, error) {
	return r.get.Execute(id)
}

func (r *Repository) List(limit, offset int) ([]entity.User, error) {
	return r.list.Execute(limit, offset)
}

func (r *Repository) Update(id int64, name, email string) error {
	return r.update.Execute(id, name, email)
}

func (r *Repository) UpdateByAdmin(id int64, name, email, status string, passwordHash *string) error {
	return r.updateByAdmin.Execute(id, name, email, status, passwordHash)
}

func (r *Repository) Delete(id int64) error {
	return r.delete.Execute(id)
}
