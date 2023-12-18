package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllMessageTypeEmailRecipient(g *gin.Context) {
	var request models.MessageTypeEmailRecipient
	g.Bind(&request)

	result, err := repository.GetAllMessageTypeEmailRecipient(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetMessageTypeEmailRecipient(g *gin.Context) {
	var request models.MessageTypeEmailRecipient
	g.Bind(&request)

	result, err := repository.GetMessageTypeEmailRecipient(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateMessageTypeEmailRecipient(g *gin.Context) {
	var request models.MessageTypeEmailRecipient
	g.Bind(&request)

	message, err := repository.CreateMessageTypeEmailRecipient(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateMessageTypeEmailRecipient(g *gin.Context) {
	var request models.MessageTypeEmailRecipient
	g.Bind(&request)

	message, err := repository.UpdateMessageTypeEmailRecipient(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteMessageTypeEmailRecipient(g *gin.Context) {
	var request models.MessageTypeEmailRecipient
	g.Bind(&request)

	message, err := repository.DeleteMessageTypeEmailRecipient(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
