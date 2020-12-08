package controller

import (
	"main/repository"
)

type User struct {
	Repository repository.Repository
}

func (u *User) GetName() string {
	return u.Repository.GetUsername()
}
