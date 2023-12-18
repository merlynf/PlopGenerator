package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllMessageTypeVariable(g *gin.Context) {
	var request models.MessageTypeVariable
	g.Bind(&request)

	result, err := repository.GetAllMessageTypeVariable(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetMessageTypeVariable(g *gin.Context) {
	var request models.MessageTypeVariable
	g.Bind(&request)

	result, err := repository.GetMessageTypeVariable(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateMessageTypeVariable(g *gin.Context) {
	var request models.MessageTypeVariable
	g.Bind(&request)

	message, err := repository.CreateMessageTypeVariable(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateMessageTypeVariable(g *gin.Context) {
	var request models.MessageTypeVariable
	g.Bind(&request)

	message, err := repository.UpdateMessageTypeVariable(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteMessageTypeVariable(g *gin.Context) {
	var request models.MessageTypeVariable
	g.Bind(&request)

	message, err := repository.DeleteMessageTypeVariable(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
