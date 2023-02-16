package database

import "gorm.io/gorm"

type AdminSettings struct {
	gorm.Model
	Password string
}

func (AdminSettings) TableName() string {
	return "Settings"
}
