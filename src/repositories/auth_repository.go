package repositories

import (
	"go-boilerplate/src/constants"

	"github.com/sarulabs/di"
	"gorm.io/gorm"
)

type AuthRepository interface {

}

type AuthRepositoryImpl struct {
	db	*gorm.DB
}

func NewAuthRepository(ioc di.Container) *AuthRepositoryImpl {
	return &AuthRepositoryImpl{
		db: ioc.Get(constants.POSTGRES).(*gorm.DB),
	}
}
