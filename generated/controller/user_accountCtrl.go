package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllUserAccount(g *gin.Context) {
	var request models.UserAccount
	g.Bind(&request)

	result, err := repository.GetAllUserAccount(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetUserAccount(g *gin.Context) {
	var request models.UserAccount
	g.Bind(&request)

	result, err := repository.GetUserAccount(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateUserAccount(g *gin.Context) {
	var request models.UserAccount
	g.Bind(&request)

	message, err := repository.CreateUserAccount(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateUserAccount(g *gin.Context) {
	var request models.UserAccount
	g.Bind(&request)

	message, err := repository.UpdateUserAccount(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteUserAccount(g *gin.Context) {
	var request models.UserAccount
	g.Bind(&request)

	message, err := repository.DeleteUserAccount(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
