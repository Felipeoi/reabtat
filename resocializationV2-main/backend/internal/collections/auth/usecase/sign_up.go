package usecase

import (
	"github.com/jeje-gab/resocializationV2/backend/internal/collections/auth"
	"golang.org/x/crypto/bcrypt"
)

type SignUp struct {
	repo auth.Repository
}

func NewSignUp(repo auth.Repository) *SignUp {
	return &SignUp{repo: repo}
}

func (u *SignUp) Execute(name, email, password, telefone string) (int64, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	resp, err := u.repo.CreateUser(name, email, string(hash), telefone)
	if err != nil {
		return 0, err
	}

	return resp, nil
}
