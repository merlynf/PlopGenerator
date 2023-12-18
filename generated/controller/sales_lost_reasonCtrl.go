package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllSalesLostReason(g *gin.Context) {
	var request models.SalesLostReason
	g.Bind(&request)

	result, err := repository.GetAllSalesLostReason(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetSalesLostReason(g *gin.Context) {
	var request models.SalesLostReason
	g.Bind(&request)

	result, err := repository.GetSalesLostReason(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateSalesLostReason(g *gin.Context) {
	var request models.SalesLostReason
	g.Bind(&request)

	message, err := repository.CreateSalesLostReason(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateSalesLostReason(g *gin.Context) {
	var request models.SalesLostReason
	g.Bind(&request)

	message, err := repository.UpdateSalesLostReason(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteSalesLostReason(g *gin.Context) {
	var request models.SalesLostReason
	g.Bind(&request)

	message, err := repository.DeleteSalesLostReason(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
