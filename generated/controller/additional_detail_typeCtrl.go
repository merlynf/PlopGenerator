package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllAdditionalDetailType(g *gin.Context) {
	var request models.AdditionalDetailType
	g.Bind(&request)

	result, err := repository.GetAllAdditionalDetailType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetAdditionalDetailType(g *gin.Context) {
	var request models.AdditionalDetailType
	g.Bind(&request)

	result, err := repository.GetAdditionalDetailType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateAdditionalDetailType(g *gin.Context) {
	var request models.AdditionalDetailType
	g.Bind(&request)

	message, err := repository.CreateAdditionalDetailType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateAdditionalDetailType(g *gin.Context) {
	var request models.AdditionalDetailType
	g.Bind(&request)

	message, err := repository.UpdateAdditionalDetailType(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteAdditionalDetailType(g *gin.Context) {
	var request models.AdditionalDetailType
	g.Bind(&request)

	message, err := repository.DeleteAdditionalDetailType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
