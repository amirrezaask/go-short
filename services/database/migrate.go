package database

import "go-short/models"

// Migrate will migrate database tables
func Migrate() {
	ORM().AutoMigrate(&models.Url{})
}
