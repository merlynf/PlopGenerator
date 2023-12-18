package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllEventLog(g *gin.Context) {
	var request models.EventLog
	g.Bind(&request)

	result, err := repository.GetAllEventLog(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetEventLog(g *gin.Context) {
	var request models.EventLog
	g.Bind(&request)

	result, err := repository.GetEventLog(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateEventLog(g *gin.Context) {
	var request models.EventLog
	g.Bind(&request)

	message, err := repository.CreateEventLog(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateEventLog(g *gin.Context) {
	var request models.EventLog
	g.Bind(&request)

	message, err := repository.UpdateEventLog(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteEventLog(g *gin.Context) {
	var request models.EventLog
	g.Bind(&request)

	message, err := repository.DeleteEventLog(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
