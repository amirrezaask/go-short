package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"go-short/config"
	"go-short/models"
)

var instance *gorm.DB

//ORM returs an instance of database.
func ORM() *gorm.DB {
	return instance
}

//InitDatbase inits database connection using gorm.
func InitDatbase() error {
	var err error
	instance, err = gorm.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8&parseTime=True&loc=Local", config.DatabaseUsername, config.DatabasePasword, config.DatabaseHost))
	if err != nil {
		return errors.Wrap(err, "error in connecting to database")
	}
	Migrate()
	return nil
}

// Migrate will migrate database tables
func Migrate() {
	ORM().AutoMigrate(&models.Url{})
}
