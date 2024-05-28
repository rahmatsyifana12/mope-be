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

type AuthController interface {
	Login(c echo.Context) error
	Logout(c echo.Context) error
}

type AuthControllerImpl struct {
	service	*services.Service
}

func NewAuthController(ioc di.Container) *AuthControllerImpl {
	return &AuthControllerImpl{
		service: ioc.Get(constants.SERVICE).(*services.Service),
	}
}

func (t *AuthControllerImpl) Login(c echo.Context) (err error) {
	var (
		params	dtos.LoginRequest
	)

	if err = c.Bind(&params); err != nil {
		return responses.NewError().
			WithCode(http.StatusBadRequest).
			WithError(err).
			WithMessage("Failed to bind parameters").
			SendErrorResponse(c)
	}

	res, err := t.service.Auth.Login(c, params)
	return responses.New().
		WithError(err).
		WithSuccessCode(http.StatusOK).
		WithMessage("Successfully logged in").
		WithData(res).
		Send(c)
}

func (t *AuthControllerImpl) Logout(c echo.Context) (err error) {
	authClaims, err := helpers.GetAuthClaims(c)
	if err != nil {
		return responses.NewError().
			WithError(err).
			WithCode(http.StatusInternalServerError).
			WithMessage("Failed to get auth claims")
	}

	err = t.service.Auth.Logout(c,  authClaims)
	return responses.New().
		WithError(err).
		WithSuccessCode(http.StatusOK).
		WithMessage("Successfully logged out").
		Send(c)
}