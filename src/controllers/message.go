package controllers

import (
	"errors"
	"go-chat-service/src/db"
	"go-chat-service/src/dto"
	"go-chat-service/src/models"
	"go-chat-service/src/persistent"
	"go-chat-service/src/util"
	"net/http"

	"gorm.io/gorm"
)

func AddMessage(messageDTO *dto.MessageDTO) (*util.HttpResponse, error) {
	var retrievedApp models.Application

	res := db.DbInstance.Model(&models.Application{}).Where(
		models.Application{Token: messageDTO.ApplicationToken},
	).First(&retrievedApp)

	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return &util.HttpResponse{
			StatusCode: http.StatusNotFound,
			Data:       util.HttpErrors{Errors: []string{"can't find an application with given token"}},
		}, nil
	} else if res.Error != nil {
		return nil, res.Error
	}

	var retrievedChat models.Chat

	res = db.DbInstance.Model(&models.Chat{}).Where(
		models.Chat{ApplicationId: retrievedApp.ID, PerAppId: messageDTO.PerAppId},
	).First(&retrievedChat)

	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return &util.HttpResponse{
			StatusCode: http.StatusNotFound,
			Data:       util.HttpErrors{Errors: []string{"can't find a chat with given chat number"}},
		}, nil
	} else if res.Error != nil {
		return nil, res.Error
	}

	message := &models.Message{
		ChatId: retrievedChat.ID,
		Body:   messageDTO.Body,
	}

	go persistent.PersisteMessage(message)

	return &util.HttpResponse{
		StatusCode: http.StatusCreated,
		Data:       &dto.MessageCreationResponse{NumberOfMessages: retrievedChat.MessageCount + 1},
	}, nil

}
