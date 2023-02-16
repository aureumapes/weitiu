package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db, _ = gorm.Open(sqlite.Open("data/weitiu.db"), &gorm.Config{})
