package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllSalesEventType(g *gin.Context) {
	var request models.SalesEventType
	g.Bind(&request)

	result, err := repository.GetAllSalesEventType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetSalesEventType(g *gin.Context) {
	var request models.SalesEventType
	g.Bind(&request)

	result, err := repository.GetSalesEventType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateSalesEventType(g *gin.Context) {
	var request models.SalesEventType
	g.Bind(&request)

	message, err := repository.CreateSalesEventType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateSalesEventType(g *gin.Context) {
	var request models.SalesEventType
	g.Bind(&request)

	message, err := repository.UpdateSalesEventType(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteSalesEventType(g *gin.Context) {
	var request models.SalesEventType
	g.Bind(&request)

	message, err := repository.DeleteSalesEventType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
