package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllLanguageSkillLevel(g *gin.Context) {
	var request models.LanguageSkillLevel
	g.Bind(&request)

	result, err := repository.GetAllLanguageSkillLevel(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetLanguageSkillLevel(g *gin.Context) {
	var request models.LanguageSkillLevel
	g.Bind(&request)

	result, err := repository.GetLanguageSkillLevel(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateLanguageSkillLevel(g *gin.Context) {
	var request models.LanguageSkillLevel
	g.Bind(&request)

	message, err := repository.CreateLanguageSkillLevel(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateLanguageSkillLevel(g *gin.Context) {
	var request models.LanguageSkillLevel
	g.Bind(&request)

	message, err := repository.UpdateLanguageSkillLevel(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteLanguageSkillLevel(g *gin.Context) {
	var request models.LanguageSkillLevel
	g.Bind(&request)

	message, err := repository.DeleteLanguageSkillLevel(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
