package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllUserThrottle(g *gin.Context) {
	var request models.UserThrottle
	g.Bind(&request)

	result, err := repository.GetAllUserThrottle(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetUserThrottle(g *gin.Context) {
	var request models.UserThrottle
	g.Bind(&request)

	result, err := repository.GetUserThrottle(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateUserThrottle(g *gin.Context) {
	var request models.UserThrottle
	g.Bind(&request)

	message, err := repository.CreateUserThrottle(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateUserThrottle(g *gin.Context) {
	var request models.UserThrottle
	g.Bind(&request)

	message, err := repository.UpdateUserThrottle(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteUserThrottle(g *gin.Context) {
	var request models.UserThrottle
	g.Bind(&request)

	message, err := repository.DeleteUserThrottle(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
