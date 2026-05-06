package usecase

import (
	"strings"

	domain "github.com/jeje-gab/resocializationV2/backend/internal/collections/user"
	"golang.org/x/crypto/bcrypt"
)

type CreateUC struct{ repo domain.Repository }

func NewCreateUC(repo domain.Repository) *CreateUC { return &CreateUC{repo: repo} }

func (uc *CreateUC) Execute(name, email, password string) (int64, error) {
	// Normalização leve de email (opcional)
	email = strings.TrimSpace(strings.ToLower(email))

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}
	return uc.repo.Create(name, email, string(hash))
}
