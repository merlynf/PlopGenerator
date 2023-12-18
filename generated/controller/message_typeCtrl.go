package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllMessageType(g *gin.Context) {
	var request models.MessageType
	g.Bind(&request)

	result, err := repository.GetAllMessageType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetMessageType(g *gin.Context) {
	var request models.MessageType
	g.Bind(&request)

	result, err := repository.GetMessageType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateMessageType(g *gin.Context) {
	var request models.MessageType
	g.Bind(&request)

	message, err := repository.CreateMessageType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateMessageType(g *gin.Context) {
	var request models.MessageType
	g.Bind(&request)

	message, err := repository.UpdateMessageType(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteMessageType(g *gin.Context) {
	var request models.MessageType
	g.Bind(&request)

	message, err := repository.DeleteMessageType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
