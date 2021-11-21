package elasticsearch

import (
	"bytes"
	"encoding/json"
	"go-chat-service/src/models"
	"strconv"

	elasticsearch "github.com/elastic/go-elasticsearch/v7"
)

type Message struct {
	client *elasticsearch.Client
	index  string
}

type IndexedMessage struct {
	Id   uint   `json:"id"`
	Body string `json:"body"`
}

var message Message

func IndexMessage(messageInstance *models.Message) error {
	body := &IndexedMessage{
		Id:   messageInstance.ID,
		Body: messageInstance.Body,
	}
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(body); err != nil {
		return err
	}
	_, err := message.client.Index(
		message.index,
		&buf,
		message.client.Index.WithDocumentID(strconv.Itoa(int(messageInstance.ID))),
	)
	if err != nil {
		return err
	}
	return nil
}
