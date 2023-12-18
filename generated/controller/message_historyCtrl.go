package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllMessageHistory(g *gin.Context) {
	var request models.MessageHistory
	g.Bind(&request)

	result, err := repository.GetAllMessageHistory(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetMessageHistory(g *gin.Context) {
	var request models.MessageHistory
	g.Bind(&request)

	result, err := repository.GetMessageHistory(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateMessageHistory(g *gin.Context) {
	var request models.MessageHistory
	g.Bind(&request)

	message, err := repository.CreateMessageHistory(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateMessageHistory(g *gin.Context) {
	var request models.MessageHistory
	g.Bind(&request)

	message, err := repository.UpdateMessageHistory(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteMessageHistory(g *gin.Context) {
	var request models.MessageHistory
	g.Bind(&request)

	message, err := repository.DeleteMessageHistory(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
