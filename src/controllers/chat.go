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

func AddChat(chatDTO *dto.ChatDTO) (*util.HttpResponse, error) {
	var retrievedApp models.Application

	res := db.DbInstance.Model(&models.Application{}).Where(
		models.Application{Token: chatDTO.ApplicationToken},
	).First(&retrievedApp)

	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return &util.HttpResponse{
			StatusCode: http.StatusNotFound,
			Data:       util.HttpErrors{Errors: []string{"can't find an application with given token"}},
		}, nil
	} else if res.Error != nil {
		return nil, res.Error
	}

	chat := &models.Chat{ApplicationId: retrievedApp.ID}

	go persistent.PersisteChat(chat)

	return &util.HttpResponse{
		StatusCode: http.StatusCreated,
		Data:       &dto.ChatCreationResponse{NumberOfChats: retrievedApp.ChatCount + 1},
	}, nil

}
