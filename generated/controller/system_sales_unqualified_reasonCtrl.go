package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllSystemSalesUnqualifiedReason(g *gin.Context) {
	var request models.SystemSalesUnqualifiedReason
	g.Bind(&request)

	result, err := repository.GetAllSystemSalesUnqualifiedReason(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetSystemSalesUnqualifiedReason(g *gin.Context) {
	var request models.SystemSalesUnqualifiedReason
	g.Bind(&request)

	result, err := repository.GetSystemSalesUnqualifiedReason(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateSystemSalesUnqualifiedReason(g *gin.Context) {
	var request models.SystemSalesUnqualifiedReason
	g.Bind(&request)

	message, err := repository.CreateSystemSalesUnqualifiedReason(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateSystemSalesUnqualifiedReason(g *gin.Context) {
	var request models.SystemSalesUnqualifiedReason
	g.Bind(&request)

	message, err := repository.UpdateSystemSalesUnqualifiedReason(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteSystemSalesUnqualifiedReason(g *gin.Context) {
	var request models.SystemSalesUnqualifiedReason
	g.Bind(&request)

	message, err := repository.DeleteSystemSalesUnqualifiedReason(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
