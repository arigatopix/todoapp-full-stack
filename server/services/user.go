package services

import "server/models"

type User struct {
	ID    int
	Email string
	Password string
}

func (u *User) ExistByEmail() (bool, error) {
	return models.ExistEmail(u.Email)
}
