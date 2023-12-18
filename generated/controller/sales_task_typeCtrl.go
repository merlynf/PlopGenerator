package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllSalesTaskType(g *gin.Context) {
	var request models.SalesTaskType
	g.Bind(&request)

	result, err := repository.GetAllSalesTaskType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetSalesTaskType(g *gin.Context) {
	var request models.SalesTaskType
	g.Bind(&request)

	result, err := repository.GetSalesTaskType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateSalesTaskType(g *gin.Context) {
	var request models.SalesTaskType
	g.Bind(&request)

	message, err := repository.CreateSalesTaskType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateSalesTaskType(g *gin.Context) {
	var request models.SalesTaskType
	g.Bind(&request)

	message, err := repository.UpdateSalesTaskType(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteSalesTaskType(g *gin.Context) {
	var request models.SalesTaskType
	g.Bind(&request)

	message, err := repository.DeleteSalesTaskType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
