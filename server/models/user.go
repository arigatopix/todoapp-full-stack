package models

import "gorm.io/gorm"

type User struct {
	Model
	Email string `json:"email"`
}

func AddUser(data map[string]interface{}) (*User, error) {
	db := ConnectDB()

	user := User{
		Email: data["email"].(string),
	}

	if err := db.Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUser(id int) (*User, error) {
	db := ConnectDB()
	var user User

	err := db.First(&user, id).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &user, nil
}
