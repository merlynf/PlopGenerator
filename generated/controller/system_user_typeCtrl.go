package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllSystemUserType(g *gin.Context) {
	var request models.SystemUserType
	g.Bind(&request)

	result, err := repository.GetAllSystemUserType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetSystemUserType(g *gin.Context) {
	var request models.SystemUserType
	g.Bind(&request)

	result, err := repository.GetSystemUserType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateSystemUserType(g *gin.Context) {
	var request models.SystemUserType
	g.Bind(&request)

	message, err := repository.CreateSystemUserType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateSystemUserType(g *gin.Context) {
	var request models.SystemUserType
	g.Bind(&request)

	message, err := repository.UpdateSystemUserType(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteSystemUserType(g *gin.Context) {
	var request models.SystemUserType
	g.Bind(&request)

	message, err := repository.DeleteSystemUserType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
