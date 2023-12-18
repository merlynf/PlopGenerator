package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllLanguageSkill(g *gin.Context) {
	var request models.LanguageSkill
	g.Bind(&request)

	result, err := repository.GetAllLanguageSkill(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetLanguageSkill(g *gin.Context) {
	var request models.LanguageSkill
	g.Bind(&request)

	result, err := repository.GetLanguageSkill(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateLanguageSkill(g *gin.Context) {
	var request models.LanguageSkill
	g.Bind(&request)

	message, err := repository.CreateLanguageSkill(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateLanguageSkill(g *gin.Context) {
	var request models.LanguageSkill
	g.Bind(&request)

	message, err := repository.UpdateLanguageSkill(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteLanguageSkill(g *gin.Context) {
	var request models.LanguageSkill
	g.Bind(&request)

	message, err := repository.DeleteLanguageSkill(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
