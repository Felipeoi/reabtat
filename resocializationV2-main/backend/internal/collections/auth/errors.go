package auth

import "errors"

// Erros de autenticação consumidos pelo handler HTTP (errors.Is).
var (
	ErrInvalidLogin    = errors.New("auth: invalid login")
	ErrInactiveAccount = errors.New("auth: inactive account")
)
