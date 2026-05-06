package match

import (
	"context"

	"github.com/jeje-gab/resocializationV2/backend/internal/entity"
)

type Usecase interface {
	// FindMatches busca todos os matches possíveis para os inmates do usuário logado
	FindMatches(ctx context.Context) ([]entity.MatchResult, error)
	// FindMatchByID busca um match específico pelos IDs dos inmates
	FindMatchByID(ctx context.Context, myInmateID, matchedInmateID int) (*entity.MatchResult, error)
}
