package main

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"worker/src/web"
)

type (
	CustomValidator struct {
		validator *validator.Validate
	}
)

func main() {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/health", web.GetHealth)
	e.GET("/info", web.GetVMInfo)
	e.GET("/run", web.RunFunction)

	e.Logger.Fatal(e.Start(":8080"))
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
