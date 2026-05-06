package match

import (
	"context"

	"github.com/jeje-gab/resocializationV2/backend/internal/entity"
)

type Repository interface {
	// FindMatches busca todos os matches possíveis para os inmates do usuário
	FindMatches(ctx context.Context, userID int64) ([]entity.MatchResult, error)
	// FindMatchByID busca um match específico pelos IDs dos inmates
	FindMatchByID(ctx context.Context, userID int64, myInmateID, matchedInmateID int) (*entity.MatchResult, error)
}
