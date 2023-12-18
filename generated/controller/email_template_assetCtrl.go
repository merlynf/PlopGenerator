package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllEmailTemplateAsset(g *gin.Context) {
	var request models.EmailTemplateAsset
	g.Bind(&request)

	result, err := repository.GetAllEmailTemplateAsset(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetEmailTemplateAsset(g *gin.Context) {
	var request models.EmailTemplateAsset
	g.Bind(&request)

	result, err := repository.GetEmailTemplateAsset(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateEmailTemplateAsset(g *gin.Context) {
	var request models.EmailTemplateAsset
	g.Bind(&request)

	message, err := repository.CreateEmailTemplateAsset(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateEmailTemplateAsset(g *gin.Context) {
	var request models.EmailTemplateAsset
	g.Bind(&request)

	message, err := repository.UpdateEmailTemplateAsset(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteEmailTemplateAsset(g *gin.Context) {
	var request models.EmailTemplateAsset
	g.Bind(&request)

	message, err := repository.DeleteEmailTemplateAsset(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
