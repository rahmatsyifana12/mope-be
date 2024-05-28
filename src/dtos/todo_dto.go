package dtos

import "go-boilerplate/src/models"

type GetTodoByIDResponse struct {
	Todo models.Todo
}

type GetTodosResponse struct {
	Todos []models.Todo
}

type CreateTodoRequest struct {
	Title string `json:"title"`
	Content string `json:"content"`
}

type TodoIDParams struct {
	ID uint `param:"id" validate:"required"`
}

type UpdateTodoParams struct {
	ID		uint	`param:"id" validate:"required"`
	Title   string	`json:"title"`
	Content string	`json:"content"`
}