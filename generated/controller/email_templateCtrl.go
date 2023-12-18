package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllEmailTemplate(g *gin.Context) {
	var request models.EmailTemplate
	g.Bind(&request)

	result, err := repository.GetAllEmailTemplate(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetEmailTemplate(g *gin.Context) {
	var request models.EmailTemplate
	g.Bind(&request)

	result, err := repository.GetEmailTemplate(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateEmailTemplate(g *gin.Context) {
	var request models.EmailTemplate
	g.Bind(&request)

	message, err := repository.CreateEmailTemplate(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateEmailTemplate(g *gin.Context) {
	var request models.EmailTemplate
	g.Bind(&request)

	message, err := repository.UpdateEmailTemplate(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteEmailTemplate(g *gin.Context) {
	var request models.EmailTemplate
	g.Bind(&request)

	message, err := repository.DeleteEmailTemplate(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
