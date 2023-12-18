package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllNamePrefix(g *gin.Context) {
	var request models.NamePrefix
	g.Bind(&request)

	result, err := repository.GetAllNamePrefix(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetNamePrefix(g *gin.Context) {
	var request models.NamePrefix
	g.Bind(&request)

	result, err := repository.GetNamePrefix(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateNamePrefix(g *gin.Context) {
	var request models.NamePrefix
	g.Bind(&request)

	message, err := repository.CreateNamePrefix(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateNamePrefix(g *gin.Context) {
	var request models.NamePrefix
	g.Bind(&request)

	message, err := repository.UpdateNamePrefix(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteNamePrefix(g *gin.Context) {
	var request models.NamePrefix
	g.Bind(&request)

	message, err := repository.DeleteNamePrefix(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
