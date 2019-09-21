package main

import (
	"github.com/labstack/echo"
	"go-short/services/url"
)

func registerRoutes(e *echo.Echo) {
	e.GET("/", func(ctx echo.Context) error {
		response := map[string]string{"message": "OK"}
		return ctx.JSON(200, response)
	})
	e.GET("/go/:uri", url.Redirection)
	v1 := e.Group("/api/v1")
	v1.POST("/url", url.New)
}
