package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllSystemEmailTemplate(g *gin.Context) {
	var request models.SystemEmailTemplate
	g.Bind(&request)

	result, err := repository.GetAllSystemEmailTemplate(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetSystemEmailTemplate(g *gin.Context) {
	var request models.SystemEmailTemplate
	g.Bind(&request)

	result, err := repository.GetSystemEmailTemplate(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateSystemEmailTemplate(g *gin.Context) {
	var request models.SystemEmailTemplate
	g.Bind(&request)

	message, err := repository.CreateSystemEmailTemplate(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateSystemEmailTemplate(g *gin.Context) {
	var request models.SystemEmailTemplate
	g.Bind(&request)

	message, err := repository.UpdateSystemEmailTemplate(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteSystemEmailTemplate(g *gin.Context) {
	var request models.SystemEmailTemplate
	g.Bind(&request)

	message, err := repository.DeleteSystemEmailTemplate(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
