package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Url struct {
	gorm.Model
	Uri    string `gorm:"type:varchar(5);unique_index"`
	Target string `gorm:"type:varchar(255)"`
}
