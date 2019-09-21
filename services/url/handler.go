package url

import (
	"github.com/labstack/echo"
	"go-short/database"
	"go-short/models"
)

type fields struct {
	Url string `validate:"required,url" json:"url"`
}

type apiError struct {
	message string `json:"reason"`
}

func (a apiError) Error() string {
	return a.message
}

func newapiError(err error) apiError {
	return apiError{err.Error()}
}

func New(c echo.Context) error {
	var fields fields

	if err := c.Bind(&fields); err != nil {
		return c.JSON(400, newapiError(err))
	}

	if err := c.Validate(fields); err != nil {
		return c.JSON(400, newapiError(err))
	}

	url := models.Url{
		Target: fields.Url,
		Uri:    newUri(),
	}

	if err := database.ORM().Save(&url).Error; err != nil {
		return c.JSON(500, newapiError(err))
	}
	return c.JSON(200, url)
}

func Redirection(c echo.Context) error {
	uri := c.Param("uri")

	url := new(models.Url)

	if err := database.ORM().Where("uri = ?", uri).First(&url).Error; err != nil {
		return c.JSON(500, newapiError(err))
	}

	if url == nil {
		return c.JSON(404, apiError{"Not found."})
	}
	return c.Redirect(301, url.Target)
}
