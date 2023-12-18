package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllSalesProspectAsset(g *gin.Context) {
	var request models.SalesProspectAsset
	g.Bind(&request)

	result, err := repository.GetAllSalesProspectAsset(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetSalesProspectAsset(g *gin.Context) {
	var request models.SalesProspectAsset
	g.Bind(&request)

	result, err := repository.GetSalesProspectAsset(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateSalesProspectAsset(g *gin.Context) {
	var request models.SalesProspectAsset
	g.Bind(&request)

	message, err := repository.CreateSalesProspectAsset(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateSalesProspectAsset(g *gin.Context) {
	var request models.SalesProspectAsset
	g.Bind(&request)

	message, err := repository.UpdateSalesProspectAsset(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteSalesProspectAsset(g *gin.Context) {
	var request models.SalesProspectAsset
	g.Bind(&request)

	message, err := repository.DeleteSalesProspectAsset(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
