package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllTimeOfTheDay(g *gin.Context) {
	var request models.TimeOfTheDay
	g.Bind(&request)

	result, err := repository.GetAllTimeOfTheDay(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetTimeOfTheDay(g *gin.Context) {
	var request models.TimeOfTheDay
	g.Bind(&request)

	result, err := repository.GetTimeOfTheDay(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateTimeOfTheDay(g *gin.Context) {
	var request models.TimeOfTheDay
	g.Bind(&request)

	message, err := repository.CreateTimeOfTheDay(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateTimeOfTheDay(g *gin.Context) {
	var request models.TimeOfTheDay
	g.Bind(&request)

	message, err := repository.UpdateTimeOfTheDay(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteTimeOfTheDay(g *gin.Context) {
	var request models.TimeOfTheDay
	g.Bind(&request)

	message, err := repository.DeleteTimeOfTheDay(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
