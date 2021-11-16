package main

import "go-chat-service/src/api"

func main() {
	api.Init()
	api.Serve()
}
