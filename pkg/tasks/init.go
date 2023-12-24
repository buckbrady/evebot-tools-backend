package tasks

import (
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func AddDB(database *gorm.DB) {
	db = database
}
