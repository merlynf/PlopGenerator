package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllSystemSalesTaskType(g *gin.Context) {
	var request models.SystemSalesTaskType
	g.Bind(&request)

	result, err := repository.GetAllSystemSalesTaskType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetSystemSalesTaskType(g *gin.Context) {
	var request models.SystemSalesTaskType
	g.Bind(&request)

	result, err := repository.GetSystemSalesTaskType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateSystemSalesTaskType(g *gin.Context) {
	var request models.SystemSalesTaskType
	g.Bind(&request)

	message, err := repository.CreateSystemSalesTaskType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateSystemSalesTaskType(g *gin.Context) {
	var request models.SystemSalesTaskType
	g.Bind(&request)

	message, err := repository.UpdateSystemSalesTaskType(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteSystemSalesTaskType(g *gin.Context) {
	var request models.SystemSalesTaskType
	g.Bind(&request)

	message, err := repository.DeleteSystemSalesTaskType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
