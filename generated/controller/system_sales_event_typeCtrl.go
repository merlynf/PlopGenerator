package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllSystemSalesEventType(g *gin.Context) {
	var request models.SystemSalesEventType
	g.Bind(&request)

	result, err := repository.GetAllSystemSalesEventType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetSystemSalesEventType(g *gin.Context) {
	var request models.SystemSalesEventType
	g.Bind(&request)

	result, err := repository.GetSystemSalesEventType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateSystemSalesEventType(g *gin.Context) {
	var request models.SystemSalesEventType
	g.Bind(&request)

	message, err := repository.CreateSystemSalesEventType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateSystemSalesEventType(g *gin.Context) {
	var request models.SystemSalesEventType
	g.Bind(&request)

	message, err := repository.UpdateSystemSalesEventType(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteSystemSalesEventType(g *gin.Context) {
	var request models.SystemSalesEventType
	g.Bind(&request)

	message, err := repository.DeleteSystemSalesEventType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
