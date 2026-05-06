package usecase

import (
	"github.com/jeje-gab/resocializationV2/backend/internal/collections/auth"
	"github.com/jeje-gab/resocializationV2/backend/internal/entity"
)

type UseCase struct {
	signUp *SignUp
	login  *login
	me     *Me
}

func NewUseCase(repo auth.Repository, jwt auth.JWT) auth.Usecase {
	return &UseCase{
		signUp: NewSignUp(repo),
		login:  NewLogin(repo, jwt),
		me:     NewMe(repo),
	}
}

func (u *UseCase) Signup(name, email, password, telefone string) (int64, error) {
	return u.signUp.Execute(name, email, password, telefone)
}

func (u *UseCase) Login(email, password string) (string, error) {
	return u.login.Execute(email, password)
}

func (u *UseCase) Me(userID int64) (*entity.User, error) {
	return u.me.Execute(userID)
}
