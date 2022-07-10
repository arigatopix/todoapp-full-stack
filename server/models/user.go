package models

import "gorm.io/gorm"

type User struct {
	Model
	Email string `gorm:"index:unique" json:"email"`
	Password string `json:"password"`
}

func ExistEmail(email string) (bool, error) {
	db := ConnectDB()

	var user User

	// หา email
	if err := db.Where(User{Email: email}).First(&user).Error; err != nil {
		return false, err
	}

	if user.ID > 0 {
		return true, nil
	}

	return false, nil
}

func AddUser(data map[string]interface{}) (*User, error) {
	db := ConnectDB()

	user := User{
		Email: data["email"].(string),
		Password: data["password"].(string),
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

func GetUserByEmail(email string) (*User, error) {
	db := ConnectDB()
	var user User

	err := db.Where("email = ?", email).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}
