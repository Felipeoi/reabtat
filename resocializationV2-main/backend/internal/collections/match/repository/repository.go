package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jeje-gab/resocializationV2/backend/internal/collections/match"
	"github.com/jeje-gab/resocializationV2/backend/internal/entity"
)

type Repository struct {
	findMatches   *FindMatchesRepo
	findMatchByID *FindMatchByIDRepo
}

func NewRepository(pool *pgxpool.Pool) match.Repository {
	return &Repository{
		findMatches:   NewFindMatchesRepo(pool),
		findMatchByID: NewFindMatchByIDRepo(pool),
	}
}

func (r *Repository) FindMatches(ctx context.Context, userID int64) ([]entity.MatchResult, error) {
	return r.findMatches.Execute(ctx, userID)
}

func (r *Repository) FindMatchByID(ctx context.Context, userID int64, myInmateID, matchedInmateID int) (*entity.MatchResult, error) {
	return r.findMatchByID.Execute(ctx, userID, myInmateID, matchedInmateID)
}
