package models

import "gorm.io/gorm"

type Todo struct {
	Model
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func AddTodo(data map[string]interface{}) (*Todo, error) {
	db := ConnectDB()

	todo := Todo{
		Title:     data["title"].(string),
		Completed: data["completed"].(bool),
	}

	if err := db.Create(&todo).Error; err != nil {
		return nil, err
	}

	return &todo, nil
}

func GetTodos() (*[]Todo, error) {

	db := ConnectDB()

	var todos []Todo

	if err := db.Find(&todos).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &todos, nil
}

func GetTodo(id int) (*Todo, error) {
	db := ConnectDB()
	var todo Todo

	err := db.First(&todo, id).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &todo, nil
}

func DeleteTodo(id int) error {
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
