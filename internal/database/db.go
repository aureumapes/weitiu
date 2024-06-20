package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

var home, _ = os.UserHomeDir()

var Db, _ = gorm.Open(sqlite.Open(home+"/.weitiu/resource/weitiu.sqlite"), &gorm.Config{})
