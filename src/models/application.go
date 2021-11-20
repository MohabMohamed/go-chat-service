package models

import "time"

type Application struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Token     string
	Name      string
	ChatCount int
}
