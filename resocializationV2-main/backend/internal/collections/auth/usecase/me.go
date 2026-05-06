package usecase

import (
	"github.com/jeje-gab/resocializationV2/backend/internal/collections/auth"
	"github.com/jeje-gab/resocializationV2/backend/internal/entity"
)

type Me struct {
	repo auth.Repository
}

func NewMe(repo auth.Repository) *Me {
	return &Me{repo: repo}
}

func (m *Me) Execute(userID int64) (*entity.User, error) {
	return m.repo.GetUserByID(userID)
}
