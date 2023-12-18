package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllBusinessEntityDocument(g *gin.Context) {
	var request models.BusinessEntityDocument
	g.Bind(&request)

	result, err := repository.GetAllBusinessEntityDocument(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetBusinessEntityDocument(g *gin.Context) {
	var request models.BusinessEntityDocument
	g.Bind(&request)

	result, err := repository.GetBusinessEntityDocument(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateBusinessEntityDocument(g *gin.Context) {
	var request models.BusinessEntityDocument
	g.Bind(&request)

	message, err := repository.CreateBusinessEntityDocument(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateBusinessEntityDocument(g *gin.Context) {
	var request models.BusinessEntityDocument
	g.Bind(&request)

	message, err := repository.UpdateBusinessEntityDocument(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteBusinessEntityDocument(g *gin.Context) {
	var request models.BusinessEntityDocument
	g.Bind(&request)

	message, err := repository.DeleteBusinessEntityDocument(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
