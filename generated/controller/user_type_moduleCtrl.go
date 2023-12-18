package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllUserTypeModule(g *gin.Context) {
	var request models.UserTypeModule
	g.Bind(&request)

	result, err := repository.GetAllUserTypeModule(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetUserTypeModule(g *gin.Context) {
	var request models.UserTypeModule
	g.Bind(&request)

	result, err := repository.GetUserTypeModule(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateUserTypeModule(g *gin.Context) {
	var request models.UserTypeModule
	g.Bind(&request)

	message, err := repository.CreateUserTypeModule(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateUserTypeModule(g *gin.Context) {
	var request models.UserTypeModule
	g.Bind(&request)

	message, err := repository.UpdateUserTypeModule(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteUserTypeModule(g *gin.Context) {
	var request models.UserTypeModule
	g.Bind(&request)

	message, err := repository.DeleteUserTypeModule(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
