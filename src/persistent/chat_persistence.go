package persistent

import (
	"go-chat-service/src/db"
	"go-chat-service/src/models"
	"sync"
	"time"
)

type chatQueue struct {
	sync.Mutex
	chats []*models.Chat
}

var chatQueueInstance chatQueue

func PersisteChat(chat *models.Chat, app_id int) {
	chatQueueInstance.Lock()
	defer chatQueueInstance.Unlock()
	max_id := 0
	if len(chatQueueInstance.chats) == 0 {

		db.DbInstance.Model(&models.Chat{}).Select("MAX(per_app_id)").
			Where("application_id = ?", app_id).Group("application_id").First(&max_id)
		chat.PerAppId = max_id + 1
	}
	chatQueueInstance.chats = append(chatQueueInstance.chats, chat)
}

func PersisteChatsQueue() {
	for range time.Tick(time.Minute * 2) {
		chatQueueInstance.Lock()
		defer chatQueueInstance.Unlock()

		if len(chatQueueInstance.chats) > 0 {
			db.DbInstance.Model(&models.Chat{}).CreateInBatches(chatQueueInstance.chats, 2000)
			chatQueueInstance.chats = nil
		}
	}
}
