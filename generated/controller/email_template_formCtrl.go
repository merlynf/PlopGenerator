package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllEmailTemplateForm(g *gin.Context) {
	var request models.EmailTemplateForm
	g.Bind(&request)

	result, err := repository.GetAllEmailTemplateForm(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetEmailTemplateForm(g *gin.Context) {
	var request models.EmailTemplateForm
	g.Bind(&request)

	result, err := repository.GetEmailTemplateForm(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateEmailTemplateForm(g *gin.Context) {
	var request models.EmailTemplateForm
	g.Bind(&request)

	message, err := repository.CreateEmailTemplateForm(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateEmailTemplateForm(g *gin.Context) {
	var request models.EmailTemplateForm
	g.Bind(&request)

	message, err := repository.UpdateEmailTemplateForm(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteEmailTemplateForm(g *gin.Context) {
	var request models.EmailTemplateForm
	g.Bind(&request)

	message, err := repository.DeleteEmailTemplateForm(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
