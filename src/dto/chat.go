package dto

type ChatDTO struct {
	ApplicationToken string `json:"token"`
}

type ChatCreationResponse struct {
	NumberOfChats int `json:"number_of_chats"`
}
