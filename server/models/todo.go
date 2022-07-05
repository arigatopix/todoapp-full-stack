package models

import (
	"gorm.io/gorm"
)

type Todo struct {
	Model
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	UserId    int    `json:"user_id"`
}

func AddTodo(data map[string]interface{}) (*Todo, error) {
	db := ConnectDB()

	todo := Todo{
		Title:     data["title"].(string),
		Completed: data["completed"].(bool),
		UserId:    data["user_id"].(int),
	}

	if err := db.Create(&todo).Error; err != nil {
		return nil, err
	}

	return &todo, nil
}

func GetTodos(userId int) (*[]Todo, error) {

	db := ConnectDB()

	var todos []Todo

	if err := db.Where("user_id = ?", userId).Find(&todos).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &todos, nil
}

func GetTodo(id int, userId int) (*Todo, error) {
	db := ConnectDB()
	var todo Todo

	if err := db.Where("ID = ? AND user_id = ?", id, userId).First(&todo).Error; err != nil {
		return nil, err
	}

	return &todo, nil
}

func DeleteTodo(id, userId int) error {
	db := ConnectDB()

	if err := db.Delete(&Todo{}, id).Error; err != nil {
		return err
	}

	return nil
}

func UpdateTodo(id int, data interface{}) (*Todo, error) {
	db := ConnectDB()

	var todo Todo

	err := db.Model(&todo).Where("ID = ?", id).Updates(data).Error

	if err != nil {
		return nil, err
	}

	if err := db.First(&todo, id).Error; err != nil {
		return nil, err
	}

	return &todo, nil
}
