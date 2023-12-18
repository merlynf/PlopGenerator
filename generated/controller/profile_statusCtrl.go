package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllProfileStatus(g *gin.Context) {
	var request models.ProfileStatus
	g.Bind(&request)

	result, err := repository.GetAllProfileStatus(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetProfileStatus(g *gin.Context) {
	var request models.ProfileStatus
	g.Bind(&request)

	result, err := repository.GetProfileStatus(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateProfileStatus(g *gin.Context) {
	var request models.ProfileStatus
	g.Bind(&request)

	message, err := repository.CreateProfileStatus(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateProfileStatus(g *gin.Context) {
	var request models.ProfileStatus
	g.Bind(&request)

	message, err := repository.UpdateProfileStatus(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteProfileStatus(g *gin.Context) {
	var request models.ProfileStatus
	g.Bind(&request)

	message, err := repository.DeleteProfileStatus(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
