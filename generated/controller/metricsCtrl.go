package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllMetrics(g *gin.Context) {
	var request models.Metrics
	g.Bind(&request)

	result, err := repository.GetAllMetrics(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetMetrics(g *gin.Context) {
	var request models.Metrics
	g.Bind(&request)

	result, err := repository.GetMetrics(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateMetrics(g *gin.Context) {
	var request models.Metrics
	g.Bind(&request)

	message, err := repository.CreateMetrics(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateMetrics(g *gin.Context) {
	var request models.Metrics
	g.Bind(&request)

	message, err := repository.UpdateMetrics(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteMetrics(g *gin.Context) {
	var request models.Metrics
	g.Bind(&request)

	message, err := repository.DeleteMetrics(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
