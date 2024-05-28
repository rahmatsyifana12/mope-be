package repositories

import "github.com/sarulabs/di"

type Repository struct {
	User	UserRepository
	Auth	AuthRepository
	Todo	TodoRepository
}

func NewRepository(ioc di.Container) *Repository {
	return &Repository{
		User: NewUserRepository(ioc),
		Auth: NewAuthRepository(ioc),
		Todo: NewTodoRepository(ioc),
	}
}
