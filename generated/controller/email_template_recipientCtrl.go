package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllEmailTemplateRecipient(g *gin.Context) {
	var request models.EmailTemplateRecipient
	g.Bind(&request)

	result, err := repository.GetAllEmailTemplateRecipient(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetEmailTemplateRecipient(g *gin.Context) {
	var request models.EmailTemplateRecipient
	g.Bind(&request)

	result, err := repository.GetEmailTemplateRecipient(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateEmailTemplateRecipient(g *gin.Context) {
	var request models.EmailTemplateRecipient
	g.Bind(&request)

	message, err := repository.CreateEmailTemplateRecipient(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateEmailTemplateRecipient(g *gin.Context) {
	var request models.EmailTemplateRecipient
	g.Bind(&request)

	message, err := repository.UpdateEmailTemplateRecipient(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteEmailTemplateRecipient(g *gin.Context) {
	var request models.EmailTemplateRecipient
	g.Bind(&request)

	message, err := repository.DeleteEmailTemplateRecipient(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
