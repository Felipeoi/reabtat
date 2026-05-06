package auth

import "github.com/jeje-gab/resocializationV2/backend/internal/entity"

type JWT interface {
	Generate(userID, role string) (string, error)
}

type Usecase interface {
	Signup(name, email, password, telefone string) (int64, error) // retorna ID autoincremento
	Login(email, password string) (string, error)                 // retorna JWT
	Me(userID int64) (*entity.User, error)                        // retorna dados do usuário logado
}
