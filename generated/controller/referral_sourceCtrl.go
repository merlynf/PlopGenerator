package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllReferralSource(g *gin.Context) {
	var request models.ReferralSource
	g.Bind(&request)

	result, err := repository.GetAllReferralSource(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetReferralSource(g *gin.Context) {
	var request models.ReferralSource
	g.Bind(&request)

	result, err := repository.GetReferralSource(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateReferralSource(g *gin.Context) {
	var request models.ReferralSource
	g.Bind(&request)

	message, err := repository.CreateReferralSource(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateReferralSource(g *gin.Context) {
	var request models.ReferralSource
	g.Bind(&request)

	message, err := repository.UpdateReferralSource(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteReferralSource(g *gin.Context) {
	var request models.ReferralSource
	g.Bind(&request)

	message, err := repository.DeleteReferralSource(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
