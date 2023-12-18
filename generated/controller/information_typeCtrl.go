package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllInformationType(g *gin.Context) {
	var request models.InformationType
	g.Bind(&request)

	result, err := repository.GetAllInformationType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetInformationType(g *gin.Context) {
	var request models.InformationType
	g.Bind(&request)

	result, err := repository.GetInformationType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateInformationType(g *gin.Context) {
	var request models.InformationType
	g.Bind(&request)

	message, err := repository.CreateInformationType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateInformationType(g *gin.Context) {
	var request models.InformationType
	g.Bind(&request)

	message, err := repository.UpdateInformationType(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteInformationType(g *gin.Context) {
	var request models.InformationType
	g.Bind(&request)

	message, err := repository.DeleteInformationType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
