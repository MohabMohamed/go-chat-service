package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	PerChatId int
	ChatId    uint
	Chat      Chat
	Body      string `gorm:"type:text"`
}
