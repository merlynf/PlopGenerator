package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllRecipientType(g *gin.Context) {
	var request models.RecipientType
	g.Bind(&request)

	result, err := repository.GetAllRecipientType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetRecipientType(g *gin.Context) {
	var request models.RecipientType
	g.Bind(&request)

	result, err := repository.GetRecipientType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateRecipientType(g *gin.Context) {
	var request models.RecipientType
	g.Bind(&request)

	message, err := repository.CreateRecipientType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateRecipientType(g *gin.Context) {
	var request models.RecipientType
	g.Bind(&request)

	message, err := repository.UpdateRecipientType(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteRecipientType(g *gin.Context) {
	var request models.RecipientType
	g.Bind(&request)

	message, err := repository.DeleteRecipientType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
