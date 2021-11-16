package persistent

import (
	"go-chat-service/src/db"
	"go-chat-service/src/models"
	"sync"
	"time"
)

type messageQueue struct {
	sync.Mutex
	messages []*models.Message
}

var messageQueueInstance *messageQueue

func PersisteMessage(message *models.Message, chat_id int) {
	messageQueueInstance.Lock()
	defer messageQueueInstance.Unlock()
	max_id := 0
	if len(messageQueueInstance.messages) == 0 {

		db.DbInstance.Model(&models.Message{}).Select("MAX(per_chat_id)").
			Where("chat_id = ?", chat_id).Group("chat_id").First(&max_id)
		message.PerChatId = max_id + 1
	}
	messageQueueInstance.messages = append(messageQueueInstance.messages, message)
}

func PersisteMessagesQueue() {
	for range time.Tick(time.Minute * 2) {
		messageQueueInstance.Lock()
		defer messageQueueInstance.Unlock()

		if len(messageQueueInstance.messages) > 0 {
			db.DbInstance.Model(&models.Message{}).CreateInBatches(messageQueueInstance.messages, 2000)
			messageQueueInstance.messages = nil

		}
	}
}
