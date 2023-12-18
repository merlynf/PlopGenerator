package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllSalesTask(g *gin.Context) {
	var request models.SalesTask
	g.Bind(&request)

	result, err := repository.GetAllSalesTask(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetSalesTask(g *gin.Context) {
	var request models.SalesTask
	g.Bind(&request)

	result, err := repository.GetSalesTask(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateSalesTask(g *gin.Context) {
	var request models.SalesTask
	g.Bind(&request)

	message, err := repository.CreateSalesTask(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateSalesTask(g *gin.Context) {
	var request models.SalesTask
	g.Bind(&request)

	message, err := repository.UpdateSalesTask(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteSalesTask(g *gin.Context) {
	var request models.SalesTask
	g.Bind(&request)

	message, err := repository.DeleteSalesTask(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
