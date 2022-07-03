package services

import (
	"server/models"
)

type Task struct {
	ID       int
	Text     string
	Day      string
	Reminder bool
}

func (t *Task) Update() (*models.Task, error) {
	task := map[string]interface{}{
		"text":     t.Text,
		"day":      t.Day,
		"reminder": t.Reminder,
	}

	updatedTask, err := models.UpdateTask(t.ID, task)

	if err != nil {
		return nil, err
	}

	return updatedTask, nil
}

func (t *Task) Delete() error {
	return models.DeleteTask(t.ID)
}

func (t *Task) Add() (*models.Task, error) {
	task := map[string]interface{}{
		"text":     t.Text,
		"day":      t.Day,
		"reminder": t.Reminder,
	}

	createdTask, err := models.AddTask(task)

	if err != nil {
		return nil, err
	}

	return createdTask, nil
}

func (t *Task) Get() (*models.Task, error) {

	var task *models.Task

	task, err := models.GetTask(t.ID)

	if err != nil {
		return nil, err
	}

	return task, nil
}

func (t *Task) GetAll() ([]*models.Task, error) {

	var tasks []*models.Task

	tasks, err := models.GetTasks()
	if err != nil {
		return nil, err
	}

	return tasks, nil
}
