package usecase

import (
	"strings"

	domain "github.com/jeje-gab/resocializationV2/backend/internal/collections/user"
)

type UpdateUC struct{ repo domain.Repository }

func NewUpdateUC(repo domain.Repository) *UpdateUC { return &UpdateUC{repo: repo} }

func (uc *UpdateUC) Execute(id int64, name, email string) error {
	email = strings.TrimSpace(strings.ToLower(email))
	return uc.repo.Update(id, name, email)
}
