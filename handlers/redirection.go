package handlers

import (
	"github.com/labstack/echo"
	"go-short/models"
	"go-short/services/database"
	"go-short/services/responses"
)

func Redirection(c echo.Context) error {
	uri := c.Param("uri")

	url := new(models.Url)

	database.ORM().Where("uri = ?", uri).First(&url)

	if url == nil {
		return c.JSON(404, responses.Error("Not found."))
	}

	return c.Redirect(301, url.Target)
}
