package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllSalesTaskHistory(g *gin.Context) {
	var request models.SalesTaskHistory
	g.Bind(&request)

	result, err := repository.GetAllSalesTaskHistory(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetSalesTaskHistory(g *gin.Context) {
	var request models.SalesTaskHistory
	g.Bind(&request)

	result, err := repository.GetSalesTaskHistory(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateSalesTaskHistory(g *gin.Context) {
	var request models.SalesTaskHistory
	g.Bind(&request)

	message, err := repository.CreateSalesTaskHistory(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateSalesTaskHistory(g *gin.Context) {
	var request models.SalesTaskHistory
	g.Bind(&request)

	message, err := repository.UpdateSalesTaskHistory(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteSalesTaskHistory(g *gin.Context) {
	var request models.SalesTaskHistory
	g.Bind(&request)

	message, err := repository.DeleteSalesTaskHistory(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
