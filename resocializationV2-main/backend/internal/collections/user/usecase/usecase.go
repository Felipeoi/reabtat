package usecase

import (
	domain "github.com/jeje-gab/resocializationV2/backend/internal/collections/user"
	entity "github.com/jeje-gab/resocializationV2/backend/internal/entity"
)

type UseCase struct {
	create        *CreateUC
	get           *GetUC
	list          *ListUC
	update        *UpdateUC
	updateByAdmin *UpdateByAdmin
	delete        *DeleteUC
}

func NewUseCase(repo domain.Repository) domain.Usecase {
	return &UseCase{
		create:        NewCreateUC(repo),
		get:           NewGetUC(repo),
		list:          NewListUC(repo),
		update:        NewUpdateUC(repo),
		updateByAdmin: NewUpdateByAdmin(repo),
		delete:        NewDeleteUC(repo),
	}
}

func (u *UseCase) Create(name, email, password string) (int64, error) {
	return u.create.Execute(name, email, password)
}

func (u *UseCase) Get(id int64) (*entity.User, error) {
	return u.get.Execute(id)
}

func (u *UseCase) List(limit, offset int) ([]entity.User, error) {
	return u.list.Execute(limit, offset)
}

func (u *UseCase) Update(id int64, name, email string) error {
	return u.update.Execute(id, name, email)
}

func (u *UseCase) UpdateByAdmin(id int64, name, email, status, password string) error {
	return u.updateByAdmin.Execute(id, name, email, status, password)
}

func (u *UseCase) Delete(id int64) error {
	return u.delete.Execute(id)
}
