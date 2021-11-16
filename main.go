package main

import (
	"go-chat-service/src/api"
	"go-chat-service/src/db"
)

func errCheck(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	errCheck(db.Init())
	api.Init()
	api.Serve()
}
