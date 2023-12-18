package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllMessageRecipient(g *gin.Context) {
	var request models.MessageRecipient
	g.Bind(&request)

	result, err := repository.GetAllMessageRecipient(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetMessageRecipient(g *gin.Context) {
	var request models.MessageRecipient
	g.Bind(&request)

	result, err := repository.GetMessageRecipient(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateMessageRecipient(g *gin.Context) {
	var request models.MessageRecipient
	g.Bind(&request)

	message, err := repository.CreateMessageRecipient(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateMessageRecipient(g *gin.Context) {
	var request models.MessageRecipient
	g.Bind(&request)

	message, err := repository.UpdateMessageRecipient(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteMessageRecipient(g *gin.Context) {
	var request models.MessageRecipient
	g.Bind(&request)

	message, err := repository.DeleteMessageRecipient(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
