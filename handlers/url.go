package handlers

import (
	"github.com/labstack/echo"
	"go-short/models"
	"go-short/services/database"
	"go-short/services/random"
	"go-short/services/responses"
)

type Fields struct {
	Url string `validate:"required,url" json:"url"`
}

func Url(c echo.Context) error {
	var fields Fields

	if err := c.Bind(&fields); err != nil {
		return err
	}

	if err := c.Validate(fields); err != nil {
		return c.JSON(200, responses.Error(err.Error()))
	}

	url := models.Url{
		Target: fields.Url,
		Uri:    random.Uri(),
	}

	database.ORM().Save(&url)

	return c.JSON(200, url)
}
