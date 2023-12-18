package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllSalesUnqualifiedReason(g *gin.Context) {
	var request models.SalesUnqualifiedReason
	g.Bind(&request)

	result, err := repository.GetAllSalesUnqualifiedReason(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetSalesUnqualifiedReason(g *gin.Context) {
	var request models.SalesUnqualifiedReason
	g.Bind(&request)

	result, err := repository.GetSalesUnqualifiedReason(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateSalesUnqualifiedReason(g *gin.Context) {
	var request models.SalesUnqualifiedReason
	g.Bind(&request)

	message, err := repository.CreateSalesUnqualifiedReason(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateSalesUnqualifiedReason(g *gin.Context) {
	var request models.SalesUnqualifiedReason
	g.Bind(&request)

	message, err := repository.UpdateSalesUnqualifiedReason(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteSalesUnqualifiedReason(g *gin.Context) {
	var request models.SalesUnqualifiedReason
	g.Bind(&request)

	message, err := repository.DeleteSalesUnqualifiedReason(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
