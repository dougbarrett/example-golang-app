package guestbook

import "github.com/jinzhu/gorm"

type Entry struct {
	gorm.Model
	Title   string
	Name    string
	Email   string
	Message string `gorm:"type:longtext"`
}
