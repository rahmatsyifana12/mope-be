package services

import "github.com/sarulabs/di"

type Service struct {
	User	UserService
	Auth	AuthService
	Todo	TodoService
}

func NewService(ioc di.Container) *Service {
	return &Service{
		User: NewUserService(ioc),
		Auth: NewAuthService(ioc),
		Todo: NewTodoService(ioc),
	}
}