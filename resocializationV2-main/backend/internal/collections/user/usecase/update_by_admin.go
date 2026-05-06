package usecase

import (
	"github.com/jeje-gab/resocializationV2/backend/internal/collections/user"
	"golang.org/x/crypto/bcrypt"
)

type UpdateByAdmin struct {
	repo user.Repository
}

func NewUpdateByAdmin(repo user.Repository) *UpdateByAdmin {
	return &UpdateByAdmin{repo: repo}
}

func (u *UpdateByAdmin) Execute(id int64, name, email, status, password string) error {
	var passwordHash *string

	// Se senha foi fornecida, gera hash
	if password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		hashStr := string(hash)
		passwordHash = &hashStr
	}

	return u.repo.UpdateByAdmin(id, name, email, status, passwordHash)
}
