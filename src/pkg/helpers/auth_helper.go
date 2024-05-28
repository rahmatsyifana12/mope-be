package helpers

import (
	"errors"
	"go-boilerplate/src/constants"
	"go-boilerplate/src/dtos"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func GenerateJWTString(claims dtos.AuthClaims) (token string, err error) {
	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtSecret := []byte(os.Getenv("JWT_ACCESS_SECRET"))
	token, err = rawToken.SignedString(jwtSecret)
	if err != nil {
		err = errors.New("failed to sign JWT")
	}
	return
}

func ParseAndValidateJWT(token string) (claims dtos.AuthClaims, err error) {
	_, err = jwt.ParseWithClaims(token, &claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_ACCESS_SECRET")), nil
	})
	if err != nil {
		return
	}

	return
}

func GetAuthClaims(ctx echo.Context) (claims dtos.AuthClaims, err error) {
	claims, ok := ctx.Get(constants.AuthClaimsKey).(dtos.AuthClaims)
	if !ok {
		err = errors.New("failed to cast context value to user's claims")
	}
	return
}
