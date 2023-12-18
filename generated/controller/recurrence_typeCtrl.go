package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllRecurrenceType(g *gin.Context) {
	var request models.RecurrenceType
	g.Bind(&request)

	result, err := repository.GetAllRecurrenceType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetRecurrenceType(g *gin.Context) {
	var request models.RecurrenceType
	g.Bind(&request)

	result, err := repository.GetRecurrenceType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateRecurrenceType(g *gin.Context) {
	var request models.RecurrenceType
	g.Bind(&request)

	message, err := repository.CreateRecurrenceType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateRecurrenceType(g *gin.Context) {
	var request models.RecurrenceType
	g.Bind(&request)

	message, err := repository.UpdateRecurrenceType(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteRecurrenceType(g *gin.Context) {
	var request models.RecurrenceType
	g.Bind(&request)

	message, err := repository.DeleteRecurrenceType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
