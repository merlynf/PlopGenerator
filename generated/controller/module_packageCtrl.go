package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllModulePackage(g *gin.Context) {
	var request models.ModulePackage
	g.Bind(&request)

	result, err := repository.GetAllModulePackage(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetModulePackage(g *gin.Context) {
	var request models.ModulePackage
	g.Bind(&request)

	result, err := repository.GetModulePackage(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateModulePackage(g *gin.Context) {
	var request models.ModulePackage
	g.Bind(&request)

	message, err := repository.CreateModulePackage(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateModulePackage(g *gin.Context) {
	var request models.ModulePackage
	g.Bind(&request)

	message, err := repository.UpdateModulePackage(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteModulePackage(g *gin.Context) {
	var request models.ModulePackage
	g.Bind(&request)

	message, err := repository.DeleteModulePackage(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
