package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllSystemAsset(g *gin.Context) {
	var request models.SystemAsset
	g.Bind(&request)

	result, err := repository.GetAllSystemAsset(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetSystemAsset(g *gin.Context) {
	var request models.SystemAsset
	g.Bind(&request)

	result, err := repository.GetSystemAsset(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateSystemAsset(g *gin.Context) {
	var request models.SystemAsset
	g.Bind(&request)

	message, err := repository.CreateSystemAsset(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateSystemAsset(g *gin.Context) {
	var request models.SystemAsset
	g.Bind(&request)

	message, err := repository.UpdateSystemAsset(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteSystemAsset(g *gin.Context) {
	var request models.SystemAsset
	g.Bind(&request)

	message, err := repository.DeleteSystemAsset(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
