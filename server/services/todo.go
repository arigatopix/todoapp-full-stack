package services

import "server/models"

type Todo struct {
	ID        int
	Title     string
	Completed bool
	UserId    int
}

func (t *Todo) GetAll() (*[]models.Todo, error) {
	var todos *[]models.Todo

	todos, err := models.GetTodos(t.UserId)

	if err != nil {
		return nil, err
	}

	return todos, nil
}

func (t *Todo) Get() (*models.Todo, error) {
	var todo *models.Todo

	todo, err := models.GetTodo(t.ID, t.UserId)

	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (t *Todo) Delete() error {
	return models.DeleteTodo(t.ID, t.UserId)
}

func (t *Todo) Add() (*models.Todo, error) {
	todo := map[string]interface{}{
		"title":     t.Title,
		"completed": t.Completed,
		"user_id":   t.UserId,
	}

	created, err := models.AddTodo(todo)

	if err != nil {
		return nil, err
	}

	return created, nil
}

func (t *Todo) Update(id int) (*models.Todo, error) {
	todo := map[string]interface{}{
		"title":     t.Title,
		"completed": t.Completed,
		"user_id":   t.UserId,
	}

	updated, err := models.UpdateTodo(t.ID, todo)

	if err != nil {
		return nil, err
	}

	return updated, nil
}
