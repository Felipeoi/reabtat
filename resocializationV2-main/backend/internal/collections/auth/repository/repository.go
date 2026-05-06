package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
	domain "github.com/jeje-gab/resocializationV2/backend/internal/collections/auth"
	"github.com/jeje-gab/resocializationV2/backend/internal/entity"
)

type Repository struct {
	createUser     *CreateUserRepo
	getUserByEmail *GetUserByEmailRepo
	getUserByID    *GetUserByIDRepo
}

func NewRepository(pool *pgxpool.Pool) domain.Repository {
	return &Repository{
		createUser:     NewCreateUserRepo(pool),
		getUserByEmail: NewGetUserByEmailRepo(pool),
		getUserByID:    NewGetUserByIDRepo(pool),
	}
}

func (r *Repository) CreateUser(name, email, passwordHash, telefone string) (int64, error) {
	return r.createUser.Execute(name, email, passwordHash, telefone)
}

func (r *Repository) GetUserByEmail(email string) (*entity.UserAuth, error) {
	return r.getUserByEmail.Execute(email)
}

func (r *Repository) GetUserByID(id int64) (*entity.User, error) {
	return r.getUserByID.Execute(id)
}
