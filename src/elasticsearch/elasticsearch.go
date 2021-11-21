package elasticsearch

import (
	"log"

	elasticsearch "github.com/elastic/go-elasticsearch/v7"
)

func Init() {
	var err error
	message.client, err = elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating elasticsearch the client: %s", err)
	}

	res, err := message.client.Info()
	if err != nil {
		log.Fatalf("Error from elasticsearch getting response: %s", err)
	}

	defer res.Body.Close()
	log.Println(res)
}
