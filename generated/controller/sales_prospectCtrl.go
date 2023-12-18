package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllSalesProspect(g *gin.Context) {
	var request models.SalesProspect
	g.Bind(&request)

	result, err := repository.GetAllSalesProspect(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetSalesProspect(g *gin.Context) {
	var request models.SalesProspect
	g.Bind(&request)

	result, err := repository.GetSalesProspect(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateSalesProspect(g *gin.Context) {
	var request models.SalesProspect
	g.Bind(&request)

	message, err := repository.CreateSalesProspect(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateSalesProspect(g *gin.Context) {
	var request models.SalesProspect
	g.Bind(&request)

	message, err := repository.UpdateSalesProspect(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteSalesProspect(g *gin.Context) {
	var request models.SalesProspect
	g.Bind(&request)

	message, err := repository.DeleteSalesProspect(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
