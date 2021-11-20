package main

import (
	"go-chat-service/src/api"
	"go-chat-service/src/config"
	"go-chat-service/src/db"
	"go-chat-service/src/persistent"
)

func errCheck(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	config.Init()
	errCheck(db.Init())
	go persistent.PersisteChatsQueue()
	go persistent.PersisteMessagesQueue()
	api.Init()
	api.Serve()
}
