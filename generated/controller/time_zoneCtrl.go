package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllTimeZone(g *gin.Context) {
	var request models.TimeZone
	g.Bind(&request)

	result, err := repository.GetAllTimeZone(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetTimeZone(g *gin.Context) {
	var request models.TimeZone
	g.Bind(&request)

	result, err := repository.GetTimeZone(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateTimeZone(g *gin.Context) {
	var request models.TimeZone
	g.Bind(&request)

	message, err := repository.CreateTimeZone(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateTimeZone(g *gin.Context) {
	var request models.TimeZone
	g.Bind(&request)

	message, err := repository.UpdateTimeZone(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteTimeZone(g *gin.Context) {
	var request models.TimeZone
	g.Bind(&request)

	message, err := repository.DeleteTimeZone(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
