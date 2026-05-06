package auth

import "github.com/jeje-gab/resocializationV2/backend/internal/entity"

type Repository interface {
	CreateUser(name, email, passwordHash, telefone string) (int64, error)
	GetUserByEmail(email string) (*entity.UserAuth, error)
	GetUserByID(id int64) (*entity.User, error)
}
