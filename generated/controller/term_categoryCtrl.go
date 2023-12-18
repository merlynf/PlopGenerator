package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllTermCategory(g *gin.Context) {
	var request models.TermCategory
	g.Bind(&request)

	result, err := repository.GetAllTermCategory(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetTermCategory(g *gin.Context) {
	var request models.TermCategory
	g.Bind(&request)

	result, err := repository.GetTermCategory(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateTermCategory(g *gin.Context) {
	var request models.TermCategory
	g.Bind(&request)

	message, err := repository.CreateTermCategory(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateTermCategory(g *gin.Context) {
	var request models.TermCategory
	g.Bind(&request)

	message, err := repository.UpdateTermCategory(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteTermCategory(g *gin.Context) {
	var request models.TermCategory
	g.Bind(&request)

	message, err := repository.DeleteTermCategory(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
