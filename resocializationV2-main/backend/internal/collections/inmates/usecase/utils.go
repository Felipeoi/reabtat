package usecase

import (
	"context"
	"errors"
)

type contextKey string

const userIDKey contextKey = "userID"

var ErrUserNotAuthenticated = errors.New("user not authenticated")

// getUserIDFromContext extrai o user_id do contexto
func getUserIDFromContext(ctx context.Context) (int64, error) {
	userID, ok := ctx.Value(userIDKey).(int64)
	if !ok {
		return 0, ErrUserNotAuthenticated
	}
	return userID, nil
}

// WithUserID adiciona o user_id ao contexto
func WithUserID(ctx context.Context, userID int64) context.Context {
	return context.WithValue(ctx, userIDKey, userID)
}
