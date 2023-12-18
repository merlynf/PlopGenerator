package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllBusinessEntity(g *gin.Context) {
	var request models.BusinessEntity
	g.Bind(&request)

	result, err := repository.GetAllBusinessEntity(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetBusinessEntity(g *gin.Context) {
	var request models.BusinessEntity
	g.Bind(&request)

	result, err := repository.GetBusinessEntity(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateBusinessEntity(g *gin.Context) {
	var request models.BusinessEntity
	g.Bind(&request)

	message, err := repository.CreateBusinessEntity(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateBusinessEntity(g *gin.Context) {
	var request models.BusinessEntity
	g.Bind(&request)

	message, err := repository.UpdateBusinessEntity(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteBusinessEntity(g *gin.Context) {
	var request models.BusinessEntity
	g.Bind(&request)

	message, err := repository.DeleteBusinessEntity(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
