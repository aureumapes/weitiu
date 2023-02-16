package database

import "gorm.io/gorm"

type Entry struct {
	gorm.Model
	Name    string
	Content string
}

func (Entry) TableName() string {
	return "Entries"
}
