package usecases

import (
	"Togo/models"
	"Togo/repositories"
)

type User struct {
	Repo repositories.UserDataInterface
}

func NewUser() *User {
	return &User{Repo: repositories.NewDBC()}
}

func (this *User) FetchUser(id int) *models.User {
	return this.Repo.Fetch(id)
}

func (this *User) CreateUser(user *models.User) bool {
	return this.Repo.Insert(user)
}

func (this *User) DeleteUser(id int) error {
	return this.Repo.Delete(id)
}

func (this *User) UpdateUser(user *models.User) error {
	return this.Repo.Update(user)
}
