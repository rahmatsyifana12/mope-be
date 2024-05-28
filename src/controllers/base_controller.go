package controllers

import "github.com/sarulabs/di"

type Controller struct {	
	User	UserController
	Auth	AuthController
	Todo	TodoController
}

func NewController(ioc di.Container) *Controller {
	return &Controller{
		User: NewUserController(ioc),
		Auth: NewAuthController(ioc),
		Todo: NewTodoController(ioc),
	}
}