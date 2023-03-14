package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db, _ = gorm.Open(sqlite.Open("resource/weitiu.db"), &gorm.Config{})
