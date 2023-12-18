package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllSystemSalesLostReason(g *gin.Context) {
	var request models.SystemSalesLostReason
	g.Bind(&request)

	result, err := repository.GetAllSystemSalesLostReason(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetSystemSalesLostReason(g *gin.Context) {
	var request models.SystemSalesLostReason
	g.Bind(&request)

	result, err := repository.GetSystemSalesLostReason(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateSystemSalesLostReason(g *gin.Context) {
	var request models.SystemSalesLostReason
	g.Bind(&request)

	message, err := repository.CreateSystemSalesLostReason(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateSystemSalesLostReason(g *gin.Context) {
	var request models.SystemSalesLostReason
	g.Bind(&request)

	message, err := repository.UpdateSystemSalesLostReason(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteSystemSalesLostReason(g *gin.Context) {
	var request models.SystemSalesLostReason
	g.Bind(&request)

	message, err := repository.DeleteSystemSalesLostReason(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
