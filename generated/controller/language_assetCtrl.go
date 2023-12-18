package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllLanguageAsset(g *gin.Context) {
	var request models.LanguageAsset
	g.Bind(&request)

	result, err := repository.GetAllLanguageAsset(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetLanguageAsset(g *gin.Context) {
	var request models.LanguageAsset
	g.Bind(&request)

	result, err := repository.GetLanguageAsset(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateLanguageAsset(g *gin.Context) {
	var request models.LanguageAsset
	g.Bind(&request)

	message, err := repository.CreateLanguageAsset(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateLanguageAsset(g *gin.Context) {
	var request models.LanguageAsset
	g.Bind(&request)

	message, err := repository.UpdateLanguageAsset(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteLanguageAsset(g *gin.Context) {
	var request models.LanguageAsset
	g.Bind(&request)

	message, err := repository.DeleteLanguageAsset(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
