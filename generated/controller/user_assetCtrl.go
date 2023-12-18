package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllUserAsset(g *gin.Context) {
	var request models.UserAsset
	g.Bind(&request)

	result, err := repository.GetAllUserAsset(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetUserAsset(g *gin.Context) {
	var request models.UserAsset
	g.Bind(&request)

	result, err := repository.GetUserAsset(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateUserAsset(g *gin.Context) {
	var request models.UserAsset
	g.Bind(&request)

	message, err := repository.CreateUserAsset(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateUserAsset(g *gin.Context) {
	var request models.UserAsset
	g.Bind(&request)

	message, err := repository.UpdateUserAsset(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteUserAsset(g *gin.Context) {
	var request models.UserAsset
	g.Bind(&request)

	message, err := repository.DeleteUserAsset(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
