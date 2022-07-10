package services

import "server/models"

func (u *User) Get(id int) (*models.User, error) {
	var user *models.User

	user, err := models.GetUser(id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) Add() (*models.User, error) {
	User := map[string]interface{}{
		"email": u.Email,
		"password": u.Password,
	}

	created, err := models.AddUser(User)

	if err != nil {
		return nil, err
	}

	return created, nil
}

func (u *User) Login() (*models.User, error) {

	user, err := models.GetUserByEmail(u.Email)

	if err != nil {
		return nil, err
	}

	return user, nil
}

