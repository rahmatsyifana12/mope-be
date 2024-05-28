package controllers

import (
	"go-boilerplate/src/constants"
	"go-boilerplate/src/dtos"
	"go-boilerplate/src/pkg/helpers"
	"go-boilerplate/src/pkg/responses"
	"go-boilerplate/src/services"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sarulabs/di"
)

type TodoController interface {
	CreateTodo(c echo.Context) error
	GetTodoByID(c echo.Context) error
	GetTodos(c echo.Context) error
	UpdateTodo(c echo.Context) error
	DeleteTodo(c echo.Context) error
}

type TodoControllerImpl struct {
	service	*services.Service
}

func NewTodoController(ioc di.Container) *TodoControllerImpl {
	return &TodoControllerImpl{
		service: ioc.Get(constants.SERVICE).(*services.Service),
	}
}

func (t *TodoControllerImpl) CreateTodo(c echo.Context) (err error) {
	var (
		params	dtos.CreateTodoRequest
	)

	if err = c.Bind(&params); err != nil {
		return responses.NewError().
			WithCode(http.StatusBadRequest).
			WithError(err).
			WithMessage("Failed to bind parameters").
			SendErrorResponse(c)
	}

	claims, err := helpers.GetAuthClaims(c)
	if err != nil {
		return responses.NewError().
			WithError(err).
			WithCode(http.StatusInternalServerError).
			WithMessage("Failed to get auth claims")
	}

	err = t.service.Todo.CreateTodo(c, claims, params)
	return responses.New().
		WithError(err).
		WithSuccessCode(http.StatusCreated).
		WithMessage("Successfully created a new todo").
		Send(c)
}

func (t *TodoControllerImpl) GetTodoByID(c echo.Context) error {
	var (
		params	dtos.TodoIDParams
		err		error
	)

	if err = c.Bind(&params); err != nil {
		return responses.NewError().
			WithCode(http.StatusBadRequest).
			WithError(err).
			WithMessage("Failed to bind parameters").
			SendErrorResponse(c)
	}

	claims, err := helpers.GetAuthClaims(c)
	if err != nil {
		return responses.NewError().
			WithError(err).
			WithCode(http.StatusInternalServerError).
			WithMessage("Failed to get auth claims")
	}

	data, err := t.service.Todo.GetTodoByID(c, claims, params)
	return responses.New().
		WithError(err).
		WithSuccessCode(http.StatusOK).
		WithMessage("Successfully retrieved a todo").
		WithData(data).
		Send(c)
}

func (t *TodoControllerImpl) GetTodos(c echo.Context) error {
	claims, err := helpers.GetAuthClaims(c)
	if err != nil {
		return responses.NewError().
			WithError(err).
			WithCode(http.StatusInternalServerError).
			WithMessage("Failed to get auth claims")
	}

	data, err := t.service.Todo.GetTodos(c, claims)
	return responses.New().
		WithError(err).
		WithSuccessCode(http.StatusOK).
		WithMessage("Successfully retrieved todos").
		WithData(data).
		Send(c)
}

func (t *TodoControllerImpl) UpdateTodo(c echo.Context) error {
	var (
		params	dtos.UpdateTodoParams
		err		error
	)

	if err = c.Bind(&params); err != nil {
		return responses.NewError().
			WithCode(http.StatusBadRequest).
			WithError(err).
			WithMessage("Failed to bind parameters").
			SendErrorResponse(c)
	}

	claims, err := helpers.GetAuthClaims(c)
	if err != nil {
		return responses.NewError().
			WithError(err).
			WithCode(http.StatusInternalServerError).
			WithMessage("Failed to get auth claims")
	}

	err = t.service.Todo.UpdateTodo(c, claims, params)
	return responses.New().
		WithError(err).
		WithSuccessCode(http.StatusOK).
		WithMessage("Successfully updated a todo").
		Send(c)
}

func (t *TodoControllerImpl) DeleteTodo(c echo.Context) error {
	var (
		params	dtos.TodoIDParams
		err		error
	)

	if err = c.Bind(&params); err != nil {
		return responses.NewError().
			WithCode(http.StatusBadRequest).
			WithError(err).
			WithMessage("Failed to bind parameters").
			SendErrorResponse(c)
	}

	claims, err := helpers.GetAuthClaims(c)
	if err != nil {
		return responses.NewError().
			WithError(err).
			WithCode(http.StatusInternalServerError).
			WithMessage("Failed to get auth claims")
	}

	err = t.service.Todo.DeleteTodo(c, claims, params)
	return responses.New().
		WithError(err).
		WithSuccessCode(http.StatusOK).
		WithMessage("Successfully deleted a todo").
		Send(c)
}
