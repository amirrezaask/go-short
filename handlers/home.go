package handlers

import "github.com/labstack/echo"

func Home(c echo.Context) error {
	response := map[string]string{"message": "OK"}
	return c.JSON(200, response)
}
