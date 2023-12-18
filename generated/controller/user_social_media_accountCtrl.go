package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllUserSocialMediaAccount(g *gin.Context) {
	var request models.UserSocialMediaAccount
	g.Bind(&request)

	result, err := repository.GetAllUserSocialMediaAccount(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetUserSocialMediaAccount(g *gin.Context) {
	var request models.UserSocialMediaAccount
	g.Bind(&request)

	result, err := repository.GetUserSocialMediaAccount(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateUserSocialMediaAccount(g *gin.Context) {
	var request models.UserSocialMediaAccount
	g.Bind(&request)

	message, err := repository.CreateUserSocialMediaAccount(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateUserSocialMediaAccount(g *gin.Context) {
	var request models.UserSocialMediaAccount
	g.Bind(&request)

	message, err := repository.UpdateUserSocialMediaAccount(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteUserSocialMediaAccount(g *gin.Context) {
	var request models.UserSocialMediaAccount
	g.Bind(&request)

	message, err := repository.DeleteUserSocialMediaAccount(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
