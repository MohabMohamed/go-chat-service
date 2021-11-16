package models

import "gorm.io/gorm"

type Chat struct {
	gorm.Model
	PerAppId      int
	MessageCount  int
	ApplicationId uint
	Application   Application
}
