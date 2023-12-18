package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllEmailTemplateRecipientType(g *gin.Context) {
	var request models.EmailTemplateRecipientType
	g.Bind(&request)

	result, err := repository.GetAllEmailTemplateRecipientType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetEmailTemplateRecipientType(g *gin.Context) {
	var request models.EmailTemplateRecipientType
	g.Bind(&request)

	result, err := repository.GetEmailTemplateRecipientType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateEmailTemplateRecipientType(g *gin.Context) {
	var request models.EmailTemplateRecipientType
	g.Bind(&request)

	message, err := repository.CreateEmailTemplateRecipientType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateEmailTemplateRecipientType(g *gin.Context) {
	var request models.EmailTemplateRecipientType
	g.Bind(&request)

	message, err := repository.UpdateEmailTemplateRecipientType(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteEmailTemplateRecipientType(g *gin.Context) {
	var request models.EmailTemplateRecipientType
	g.Bind(&request)

	message, err := repository.DeleteEmailTemplateRecipientType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
