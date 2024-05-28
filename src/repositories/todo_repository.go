package repositories

import (
	"go-boilerplate/src/constants"
	"go-boilerplate/src/models"

	"github.com/labstack/echo/v4"
	"github.com/sarulabs/di"
	"gorm.io/gorm"
)

type TodoRepository interface {
	CreateTodo(c echo.Context, todo models.Todo) error
	GetTodoByID(c echo.Context, todoID uint) (todo models.Todo, err error)
	GetTodosByUserID(c echo.Context, userID uint) (todos []models.Todo, err error)
	UpdateTodo(c echo.Context, todo models.Todo) error
	DeleteTodo(c echo.Context, todo models.Todo) error
}

type TodoRepositoryImpl struct {
	db	*gorm.DB
}

func NewTodoRepository(ioc di.Container) *TodoRepositoryImpl {
	return &TodoRepositoryImpl{
		db: ioc.Get(constants.POSTGRES).(*gorm.DB),
	}
}

func (r *TodoRepositoryImpl) CreateTodo(c echo.Context, todo models.Todo) error {
	err := r.db.Create(&todo).WithContext(c.Request().Context()).Error
	return err
}

func (r *TodoRepositoryImpl) GetTodoByID(c echo.Context, todoID uint) (todo models.Todo, err error) {
	err = r.db.Where("id = ?", todoID).First(&todo).WithContext(c.Request().Context()).Error
	return
}

func (r *TodoRepositoryImpl) GetTodosByUserID(c echo.Context, userID uint) (todos []models.Todo, err error) {
	err = r.db.Where("user_id = ?", userID).Find(&todos).WithContext(c.Request().Context()).Error
	return
}

func (r *TodoRepositoryImpl) UpdateTodo(c echo.Context, todo models.Todo) error {
	err := r.db.Save(&todo).WithContext(c.Request().Context()).Error
	return err
}

func (r *TodoRepositoryImpl) DeleteTodo(c echo.Context, todo models.Todo) error {
	err := r.db.Delete(&todo).WithContext(c.Request().Context()).Error
	return err
}
