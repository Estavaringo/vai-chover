package main

import (
	"fmt"
	"os"

	"github.com/Estavaringo/vai-chover/pkg/view/components"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		component := components.Index()
		return component.Render(c.Request().Context(), c.Response().Writer)
	})
	e.Use(middleware.Logger())
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("HTTP_PORT"))))
}
