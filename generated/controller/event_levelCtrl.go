package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllEventLevel(g *gin.Context) {
	var request models.EventLevel
	g.Bind(&request)

	result, err := repository.GetAllEventLevel(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetEventLevel(g *gin.Context) {
	var request models.EventLevel
	g.Bind(&request)

	result, err := repository.GetEventLevel(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateEventLevel(g *gin.Context) {
	var request models.EventLevel
	g.Bind(&request)

	message, err := repository.CreateEventLevel(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateEventLevel(g *gin.Context) {
	var request models.EventLevel
	g.Bind(&request)

	message, err := repository.UpdateEventLevel(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteEventLevel(g *gin.Context) {
	var request models.EventLevel
	g.Bind(&request)

	message, err := repository.DeleteEventLevel(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
