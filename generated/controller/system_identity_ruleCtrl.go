package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllSystemIdentityRule(g *gin.Context) {
	var request models.SystemIdentityRule
	g.Bind(&request)

	result, err := repository.GetAllSystemIdentityRule(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetSystemIdentityRule(g *gin.Context) {
	var request models.SystemIdentityRule
	g.Bind(&request)

	result, err := repository.GetSystemIdentityRule(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateSystemIdentityRule(g *gin.Context) {
	var request models.SystemIdentityRule
	g.Bind(&request)

	message, err := repository.CreateSystemIdentityRule(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateSystemIdentityRule(g *gin.Context) {
	var request models.SystemIdentityRule
	g.Bind(&request)

	message, err := repository.UpdateSystemIdentityRule(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteSystemIdentityRule(g *gin.Context) {
	var request models.SystemIdentityRule
	g.Bind(&request)

	message, err := repository.DeleteSystemIdentityRule(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
