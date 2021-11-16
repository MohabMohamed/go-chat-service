package main

import (
	"go-chat-service/src/api"
	"go-chat-service/src/db"
	"go-chat-service/src/persistent"
)

func errCheck(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	errCheck(db.Init())
	persistent.PersisteChatsQueue()
	persistent.PersisteMessagesQueue()
	api.Init()
	api.Serve()
}
