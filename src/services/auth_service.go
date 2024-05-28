package services

import (
	"go-boilerplate/src/constants"
	"go-boilerplate/src/dtos"
	"go-boilerplate/src/pkg/helpers"
	"go-boilerplate/src/pkg/responses"
	"go-boilerplate/src/repositories"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/sarulabs/di"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService interface {
	Login(c echo.Context, params dtos.LoginRequest) (dtos.LoginResponse, error)
	Logout(c echo.Context, authClaims dtos.AuthClaims) error
}

type AuthServiceImpl struct {
	repository	*repositories.Repository
}

func NewAuthService(ioc di.Container) *AuthServiceImpl {
	return &AuthServiceImpl{
		repository: ioc.Get(constants.REPOSITORY).(*repositories.Repository),
	}
}

func (s *AuthServiceImpl) Login(c echo.Context, params dtos.LoginRequest) (res dtos.LoginResponse, err error) {
	user, err := s.repository.User.GetUserByUsername(c, params.Username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = responses.NewError().
				WithError(err).
				WithCode(http.StatusBadRequest).
				WithMessage("Cannot find user with the given username")
			return
		}
		err = responses.NewError().
			WithError(err).
			WithCode(http.StatusInternalServerError).
			WithMessage("Error while retrieving user from database")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.Password))
	if err != nil {
		err = responses.NewError().
			WithError(err).
			WithCode(http.StatusBadRequest).
			WithMessage("Incorrect password")
		return
	}

	tokenExpireDuration := (time.Hour * 24)
	currentTime := time.Now()

	token, err := helpers.GenerateJWTString(dtos.AuthClaims{
		UserID:       user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(currentTime.Add(tokenExpireDuration)),
			IssuedAt:  jwt.NewNumericDate(currentTime),
		},
	})
	if err != nil {
		err = responses.NewError().
			WithError(err).
			WithCode(http.StatusInternalServerError).
			WithMessage("Error while generating JWT")
		return
	}

	res.AccessToken = token
	return
}

func (s *AuthServiceImpl) Logout(c echo.Context, authClaims dtos.AuthClaims) (err error) {
	user, err := s.repository.User.GetUserByID(c, authClaims.UserID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			err = responses.NewError().
				WithError(err).
				WithCode(http.StatusBadRequest).
				WithMessage("Cannot find user with the given id")
			return
		}
		err = responses.NewError().
			WithError(err).
			WithCode(http.StatusInternalServerError).
			WithMessage("Error while retrieving user from database")
		return
	}

	err = s.repository.User.UpdateUser(c, user)
	if err != nil {
		err = responses.NewError().
			WithError(err).
			WithCode(http.StatusInternalServerError).
			WithMessage("Error while updating users access token into database")
		return
	}

	return
}