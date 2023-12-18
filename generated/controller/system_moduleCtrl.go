package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllSystemModule(g *gin.Context) {
	var request models.SystemModule
	g.Bind(&request)

	result, err := repository.GetAllSystemModule(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetSystemModule(g *gin.Context) {
	var request models.SystemModule
	g.Bind(&request)

	result, err := repository.GetSystemModule(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateSystemModule(g *gin.Context) {
	var request models.SystemModule
	g.Bind(&request)

	message, err := repository.CreateSystemModule(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateSystemModule(g *gin.Context) {
	var request models.SystemModule
	g.Bind(&request)

	message, err := repository.UpdateSystemModule(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteSystemModule(g *gin.Context) {
	var request models.SystemModule
	g.Bind(&request)

	message, err := repository.DeleteSystemModule(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
