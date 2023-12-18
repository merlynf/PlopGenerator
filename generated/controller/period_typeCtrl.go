package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllPeriodType(g *gin.Context) {
	var request models.PeriodType
	g.Bind(&request)

	result, err := repository.GetAllPeriodType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetPeriodType(g *gin.Context) {
	var request models.PeriodType
	g.Bind(&request)

	result, err := repository.GetPeriodType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreatePeriodType(g *gin.Context) {
	var request models.PeriodType
	g.Bind(&request)

	message, err := repository.CreatePeriodType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdatePeriodType(g *gin.Context) {
	var request models.PeriodType
	g.Bind(&request)

	message, err := repository.UpdatePeriodType(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeletePeriodType(g *gin.Context) {
	var request models.PeriodType
	g.Bind(&request)

	message, err := repository.DeletePeriodType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
