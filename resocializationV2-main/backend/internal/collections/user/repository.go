package user

import "github.com/jeje-gab/resocializationV2/backend/internal/entity"

type Repository interface {
	Create(name, email, password string) (int64, error) // opcional: criar já com senha
	Get(id int64) (*entity.User, error)
	List(limit, offset int) ([]entity.User, error)
	Update(id int64, name, email string) error
	UpdateByAdmin(id int64, name, email, status string, passwordHash *string) error
	Delete(id int64) error
}
