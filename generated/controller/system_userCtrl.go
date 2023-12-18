package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllSystemUser(g *gin.Context) {
	var request models.SystemUser
	g.Bind(&request)

	result, err := repository.GetAllSystemUser(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetSystemUser(g *gin.Context) {
	var request models.SystemUser
	g.Bind(&request)

	result, err := repository.GetSystemUser(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateSystemUser(g *gin.Context) {
	var request models.SystemUser
	g.Bind(&request)

	message, err := repository.CreateSystemUser(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateSystemUser(g *gin.Context) {
	var request models.SystemUser
	g.Bind(&request)

	message, err := repository.UpdateSystemUser(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteSystemUser(g *gin.Context) {
	var request models.SystemUser
	g.Bind(&request)

	message, err := repository.DeleteSystemUser(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
