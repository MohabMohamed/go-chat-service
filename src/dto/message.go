package dto

type MessageDTO struct {
	ApplicationToken string `json:"token"`
	PerAppId         int    `json:"chat_num"`
	Body             string `json:"body" validate:"required,min=1"`
}

type MessageCreationResponse struct {
	NumberOfMessages int `json:"number_of_messages"`
}
