package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllSalesActivityStatus(g *gin.Context) {
	var request models.SalesActivityStatus
	g.Bind(&request)

	result, err := repository.GetAllSalesActivityStatus(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetSalesActivityStatus(g *gin.Context) {
	var request models.SalesActivityStatus
	g.Bind(&request)

	result, err := repository.GetSalesActivityStatus(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateSalesActivityStatus(g *gin.Context) {
	var request models.SalesActivityStatus
	g.Bind(&request)

	message, err := repository.CreateSalesActivityStatus(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateSalesActivityStatus(g *gin.Context) {
	var request models.SalesActivityStatus
	g.Bind(&request)

	message, err := repository.UpdateSalesActivityStatus(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteSalesActivityStatus(g *gin.Context) {
	var request models.SalesActivityStatus
	g.Bind(&request)

	message, err := repository.DeleteSalesActivityStatus(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
