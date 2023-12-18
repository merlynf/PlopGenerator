package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllSalesProspectEvent(g *gin.Context) {
	var request models.SalesProspectEvent
	g.Bind(&request)

	result, err := repository.GetAllSalesProspectEvent(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetSalesProspectEvent(g *gin.Context) {
	var request models.SalesProspectEvent
	g.Bind(&request)

	result, err := repository.GetSalesProspectEvent(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateSalesProspectEvent(g *gin.Context) {
	var request models.SalesProspectEvent
	g.Bind(&request)

	message, err := repository.CreateSalesProspectEvent(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateSalesProspectEvent(g *gin.Context) {
	var request models.SalesProspectEvent
	g.Bind(&request)

	message, err := repository.UpdateSalesProspectEvent(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteSalesProspectEvent(g *gin.Context) {
	var request models.SalesProspectEvent
	g.Bind(&request)

	message, err := repository.DeleteSalesProspectEvent(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
