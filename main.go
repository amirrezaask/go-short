package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go-short/handlers"
	"go-short/services/database"
	"go-short/services/http"
	"go-short/services/validator"
	"os"
)

// main
func main() {
	configure()

	logFile, err := os.OpenFile("storage/app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	e := echo.New()
	e.Validator = validator.New()
	e.Logger.SetOutput(logFile)
	e.HTTPErrorHandler = http.ErrorHandler

	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, // 1 KB
	}))

	e.GET("/", handlers.Home)
	e.GET("/go/:uri", handlers.Redirection)

	v1 := e.Group("/api/v1")
	v1.POST("/url", handlers.Url)

	e.Logger.Fatal(e.Start(":" + port()))
}

// port will return port that app listen to
func port() string {
	if port := os.Getenv("PORT"); port != "" {
		return port
	}

	return "8080"
}

// configure will configure(bootstrap) the app
func configure() {
	database.Migrate()
}
