package middlewares

import (
	"go-boilerplate/src/pkg/responses"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

func CustomErrorHandler() echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		customErr := responses.NewError().
			WithCode(http.StatusInternalServerError).
			WithError(err).
			WithMessage("Unhandled error")

		log.Error().Err(customErr).Msg("Unhandled error")

		customErr.SendErrorResponse(c)
	}
}
