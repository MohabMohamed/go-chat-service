package persistent

import (
	"go-chat-service/src/config"
	"go-chat-service/src/db"
	"go-chat-service/src/models"
	"sync"
	"time"

	"go-chat-service/src/elasticsearch"
)

type messageQueue struct {
	sync.Mutex
	messages []models.Message
}

var messageQueueInstance messageQueue

func PersisteMessage(message *models.Message) {
	messageQueueInstance.Lock()
	defer messageQueueInstance.Unlock()
	messageQueueInstance.messages = append(messageQueueInstance.messages, *message)
}

func PersisteMessagesQueue() {
	for range time.Tick(config.TimeInSecToPersist) {
		messageQueueInstance.Lock()

		if len(messageQueueInstance.messages) > 0 {
			db.DbInstance.Model(&models.Message{}).CreateInBatches(messageQueueInstance.messages, 2000)
			for _, message := range messageQueueInstance.messages {
				elasticsearch.IndexMessage(&message)
			}
			messageQueueInstance.messages = nil

		}
		messageQueueInstance.Unlock()
	}
}
