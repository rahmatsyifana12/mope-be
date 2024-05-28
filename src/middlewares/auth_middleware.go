package middlewares

import (
	"go-boilerplate/src/constants"
	"go-boilerplate/src/pkg/helpers"
	"go-boilerplate/src/pkg/responses"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeaderList, ok := c.Request().Header["Authorization"]
		if !ok || len(authHeaderList) == 0 {
			return responses.NewError().
				WithError(constants.ERR_NOT_LOGGED_IN).
				WithCode(http.StatusUnauthorized).
				WithMessage("You don't have the permission.")
		}

		authHeader := authHeaderList[0]
		bearerPrefix := "Bearer "

		if !strings.HasPrefix(authHeader, bearerPrefix) {
			return responses.NewError().
				WithError(constants.ERR_NOT_LOGGED_IN).
				WithCode(http.StatusUnauthorized).
				WithMessage("Invalid authorization header.")
		}

		token := strings.Replace(authHeader, bearerPrefix, "", 1)
		claims, err := helpers.ParseAndValidateJWT(token)
		if err != nil {
			return responses.NewError().
				WithError(err).
				WithCode(http.StatusUnauthorized).
				WithMessage("Invalid authorization header.")
		}

		c.Set(constants.AuthClaimsKey, claims)
		c.Set(constants.AccessToken, token)

		err = next(c)
		return err
	}
}