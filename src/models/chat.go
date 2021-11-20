package models

import "time"

type Chat struct {
	ID            uint `gorm:"primaryKey"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	PerAppId      int
	MessageCount  int
	ApplicationId uint
	Application   Application
}
