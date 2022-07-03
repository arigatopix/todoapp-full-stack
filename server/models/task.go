package models

import (
	"gorm.io/gorm"
)

type Task struct {
	Model
	Text     string `json:"text"`
	Day      string `json:"day"`
	Reminder bool   `json:"reminder"`
}

func UpdateTask(id int, data interface{}) (*Task, error) {
	var task Task

	err := db.Model(&task).Where("ID = ?", id).Updates(data).Error

	if err != nil {
		return nil, err
	}

	// return updated
	if err := db.First(&task, id).Error; err != nil {
		return nil, err
	}

	return &task, nil
}

func DeleteTask(id int) error {
	if err := db.Delete(&Task{}, id).Error; err != nil {
		return err
	}

	return nil
}

func AddTask(data map[string]interface{}) (*Task, error) {

	task := Task{
		Text:     data["text"].(string),
		Day:      data["day"].(string),
		Reminder: data["reminder"].(bool),
	}

	if err := db.Create(&task).Error; err != nil {
		return nil, err
	}

	return &task, nil
}

func GetTask(id int) (*Task, error) {
	var task Task

	err := db.First(&task, id).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &task, nil
}

// List all Tasks
func GetTasks() ([]*Task, error) {

	var tasks []*Task

	err := db.Find(&tasks).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return tasks, nil
}
