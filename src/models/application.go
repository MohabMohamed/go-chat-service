package models

import "gorm.io/gorm"

type Application struct {
	gorm.Model
	Token     string
	Name      string
	ChatCount int
}
