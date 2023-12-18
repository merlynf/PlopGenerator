package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllSalesProspectRevenue(g *gin.Context) {
	var request models.SalesProspectRevenue
	g.Bind(&request)

	result, err := repository.GetAllSalesProspectRevenue(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetSalesProspectRevenue(g *gin.Context) {
	var request models.SalesProspectRevenue
	g.Bind(&request)

	result, err := repository.GetSalesProspectRevenue(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateSalesProspectRevenue(g *gin.Context) {
	var request models.SalesProspectRevenue
	g.Bind(&request)

	message, err := repository.CreateSalesProspectRevenue(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateSalesProspectRevenue(g *gin.Context) {
	var request models.SalesProspectRevenue
	g.Bind(&request)

	message, err := repository.UpdateSalesProspectRevenue(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteSalesProspectRevenue(g *gin.Context) {
	var request models.SalesProspectRevenue
	g.Bind(&request)

	message, err := repository.DeleteSalesProspectRevenue(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
