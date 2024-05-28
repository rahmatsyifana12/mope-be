package main

import (
	"fmt"
	"go-boilerplate/src/logger"
	"go-boilerplate/src/middlewares"
	"os"
	"strings"

	echo_middlewares "github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
)

func main() {
    e := echo.New()

	if err := logger.SetupLogger(); err != nil {
		panic("Failed to setup logger with error: " + err.Error())
	}

	e.Use(echo_middlewares.CORSWithConfig(echo_middlewares.CORSConfig{
		AllowOrigins: strings.Split(os.Getenv("CORS_ALLOWED_ORIGINS"), ","),
	}))
	e.Use(middlewares.Log)

	module := Module{}
	module.New(e)
    
    port, found := os.LookupEnv("PORT")
	if !found {
		port = "5000"
	}

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}