package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllSalesTaskAsset(g *gin.Context) {
	var request models.SalesTaskAsset
	g.Bind(&request)

	result, err := repository.GetAllSalesTaskAsset(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetSalesTaskAsset(g *gin.Context) {
	var request models.SalesTaskAsset
	g.Bind(&request)

	result, err := repository.GetSalesTaskAsset(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateSalesTaskAsset(g *gin.Context) {
	var request models.SalesTaskAsset
	g.Bind(&request)

	message, err := repository.CreateSalesTaskAsset(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateSalesTaskAsset(g *gin.Context) {
	var request models.SalesTaskAsset
	g.Bind(&request)

	message, err := repository.UpdateSalesTaskAsset(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteSalesTaskAsset(g *gin.Context) {
	var request models.SalesTaskAsset
	g.Bind(&request)

	message, err := repository.DeleteSalesTaskAsset(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
