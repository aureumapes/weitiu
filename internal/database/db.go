package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db, _ = gorm.Open(sqlite.Open("~/.waitiu/resource/weitiu.db"), &gorm.Config{})
