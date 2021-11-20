package persistent

import (
	"go-chat-service/src/config"
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

func PersisteChat(chat *models.Chat) {
	chatQueueInstance.Lock()
	defer chatQueueInstance.Unlock()
	chatQueueInstance.chats = append(chatQueueInstance.chats, chat)
}

func PersisteChatsQueue() {
	for range time.Tick(config.TimeInSecToPersist) {
		chatQueueInstance.Lock()

		if len(chatQueueInstance.chats) > 0 {
			db.DbInstance.Model(&models.Chat{}).CreateInBatches(chatQueueInstance.chats, 2000)
			chatQueueInstance.chats = nil
		}
		chatQueueInstance.Unlock()
	}
}
