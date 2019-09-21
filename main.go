package main

import (
	"log"
	"net/http"
	"os"

	"go-short/config"
	"go-short/database"
	"go-short/services/validator"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// main
func main() {
	e := echo.New()
	//configure echo app
	err := configure(e)
	if err != nil {
		log.Fatal(err)
	}
	//initialize database connection
	err = database.InitDatbase()
	if err != nil {
		log.Fatal(err)
	}

	e.Logger.Fatal(e.Start(":" + config.Port))
}

// configure will configure(bootstrap) the app
func configure(e *echo.Echo) error {
	logFile, err := os.OpenFile("storage/app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	e.Validator = validator.New()
	e.Logger.SetOutput(logFile)
	e.HTTPErrorHandler = errorHandler

	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, // 1 KB
	}))
	registerRoutes(e)
	return nil
}

func errorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}

	if err := c.JSON(code, map[string]string{
		"reason": err.Error(),
	}); err != nil {
		//never panic
		c.Logger().Error(err)
		return
	}
	c.Logger().Error(err)
}
