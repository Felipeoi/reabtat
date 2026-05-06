package usecase

import (
	"context"

	"github.com/jeje-gab/resocializationV2/backend/internal/collections/match"
	"github.com/jeje-gab/resocializationV2/backend/internal/entity"
)

type UseCase struct {
	repo match.Repository
}

func NewUseCase(repo match.Repository) match.Usecase {
	return &UseCase{
		repo: repo,
	}
}

func (u *UseCase) FindMatches(ctx context.Context) ([]entity.MatchResult, error) {
	userID, err := getUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	return u.repo.FindMatches(ctx, userID)
}

func (u *UseCase) FindMatchByID(ctx context.Context, myInmateID, matchedInmateID int) (*entity.MatchResult, error) {
	userID, err := getUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	return u.repo.FindMatchByID(ctx, userID, myInmateID, matchedInmateID)
}
