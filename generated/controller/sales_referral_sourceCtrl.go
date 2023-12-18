package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllSalesReferralSource(g *gin.Context) {
	var request models.SalesReferralSource
	g.Bind(&request)

	result, err := repository.GetAllSalesReferralSource(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetSalesReferralSource(g *gin.Context) {
	var request models.SalesReferralSource
	g.Bind(&request)

	result, err := repository.GetSalesReferralSource(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateSalesReferralSource(g *gin.Context) {
	var request models.SalesReferralSource
	g.Bind(&request)

	message, err := repository.CreateSalesReferralSource(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateSalesReferralSource(g *gin.Context) {
	var request models.SalesReferralSource
	g.Bind(&request)

	message, err := repository.UpdateSalesReferralSource(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteSalesReferralSource(g *gin.Context) {
	var request models.SalesReferralSource
	g.Bind(&request)

	message, err := repository.DeleteSalesReferralSource(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
