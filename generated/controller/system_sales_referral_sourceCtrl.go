package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllSystemSalesReferralSource(g *gin.Context) {
	var request models.SystemSalesReferralSource
	g.Bind(&request)

	result, err := repository.GetAllSystemSalesReferralSource(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetSystemSalesReferralSource(g *gin.Context) {
	var request models.SystemSalesReferralSource
	g.Bind(&request)

	result, err := repository.GetSystemSalesReferralSource(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateSystemSalesReferralSource(g *gin.Context) {
	var request models.SystemSalesReferralSource
	g.Bind(&request)

	message, err := repository.CreateSystemSalesReferralSource(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateSystemSalesReferralSource(g *gin.Context) {
	var request models.SystemSalesReferralSource
	g.Bind(&request)

	message, err := repository.UpdateSystemSalesReferralSource(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteSystemSalesReferralSource(g *gin.Context) {
	var request models.SystemSalesReferralSource
	g.Bind(&request)

	message, err := repository.DeleteSystemSalesReferralSource(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
