package usecase

import (
	"errors"
	"strconv"

	auth2 "github.com/jeje-gab/resocializationV2/backend/internal/collections/auth"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

type login struct {
	repo auth2.Repository
	jwt  auth2.JWT
}

func NewLogin(repo auth2.Repository, jwt auth2.JWT) *login {
	return &login{
		repo: repo,
		jwt:  jwt,
	}
}

func (u *login) Execute(email, password string) (string, error) {
	user, err := u.repo.GetUserByEmail(email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", auth2.ErrInvalidLogin
		}
		return "", err
	}
	if user == nil {
		return "", auth2.ErrInvalidLogin
	}
	if bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)) != nil {
		return "", auth2.ErrInvalidLogin
	}
	if user.Status != "ativo" {
		return "", auth2.ErrInactiveAccount
	}
	// jwt espera string; convertemos o ID int64 para string e incluímos o role
	return u.jwt.Generate(strconv.FormatInt(user.ID, 10), user.Role)
}
