package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllPasswordHistory(g *gin.Context) {
	var request models.PasswordHistory
	g.Bind(&request)

	result, err := repository.GetAllPasswordHistory(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetPasswordHistory(g *gin.Context) {
	var request models.PasswordHistory
	g.Bind(&request)

	result, err := repository.GetPasswordHistory(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreatePasswordHistory(g *gin.Context) {
	var request models.PasswordHistory
	g.Bind(&request)

	message, err := repository.CreatePasswordHistory(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdatePasswordHistory(g *gin.Context) {
	var request models.PasswordHistory
	g.Bind(&request)

	message, err := repository.UpdatePasswordHistory(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeletePasswordHistory(g *gin.Context) {
	var request models.PasswordHistory
	g.Bind(&request)

	message, err := repository.DeletePasswordHistory(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
