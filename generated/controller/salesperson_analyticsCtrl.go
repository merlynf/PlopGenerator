package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllSalespersonAnalytics(g *gin.Context) {
	var request models.SalespersonAnalytics
	g.Bind(&request)

	result, err := repository.GetAllSalespersonAnalytics(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetSalespersonAnalytics(g *gin.Context) {
	var request models.SalespersonAnalytics
	g.Bind(&request)

	result, err := repository.GetSalespersonAnalytics(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateSalespersonAnalytics(g *gin.Context) {
	var request models.SalespersonAnalytics
	g.Bind(&request)

	message, err := repository.CreateSalespersonAnalytics(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateSalespersonAnalytics(g *gin.Context) {
	var request models.SalespersonAnalytics
	g.Bind(&request)

	message, err := repository.UpdateSalespersonAnalytics(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteSalespersonAnalytics(g *gin.Context) {
	var request models.SalespersonAnalytics
	g.Bind(&request)

	message, err := repository.DeleteSalespersonAnalytics(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
