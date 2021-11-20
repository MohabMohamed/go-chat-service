package models

import "time"

type Message struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	PerChatId int
	ChatId    uint
	Chat      Chat
	Body      string `gorm:"type:text"`
}
