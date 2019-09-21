package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"os"
	"sync"
)

var (
	once     sync.Once
	instance *gorm.DB
)

func ORM() *gorm.DB {
	once.Do(func() {
		n := os.Getenv("DB_DATABASE")
		u := os.Getenv("DB_USERNAME")
		p := os.Getenv("DB_PASSWORD")
		h := os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT")

		url := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8&parseTime=True&loc=Local", u, p, h, n)

		var err error

		instance, err = gorm.Open("mysql", url)
		if err != nil {
			panic(err)
		}
	})

	return instance
}
