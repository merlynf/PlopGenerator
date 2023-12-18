package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllModuleAsset(g *gin.Context) {
	var request models.ModuleAsset
	g.Bind(&request)

	result, err := repository.GetAllModuleAsset(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetModuleAsset(g *gin.Context) {
	var request models.ModuleAsset
	g.Bind(&request)

	result, err := repository.GetModuleAsset(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateModuleAsset(g *gin.Context) {
	var request models.ModuleAsset
	g.Bind(&request)

	message, err := repository.CreateModuleAsset(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateModuleAsset(g *gin.Context) {
	var request models.ModuleAsset
	g.Bind(&request)

	message, err := repository.UpdateModuleAsset(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteModuleAsset(g *gin.Context) {
	var request models.ModuleAsset
	g.Bind(&request)

	message, err := repository.DeleteModuleAsset(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
