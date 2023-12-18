package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllCustomerAsset(g *gin.Context) {
	var request models.CustomerAsset
	g.Bind(&request)

	result, err := repository.GetAllCustomerAsset(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetCustomerAsset(g *gin.Context) {
	var request models.CustomerAsset
	g.Bind(&request)

	result, err := repository.GetCustomerAsset(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateCustomerAsset(g *gin.Context) {
	var request models.CustomerAsset
	g.Bind(&request)

	message, err := repository.CreateCustomerAsset(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateCustomerAsset(g *gin.Context) {
	var request models.CustomerAsset
	g.Bind(&request)

	message, err := repository.UpdateCustomerAsset(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteCustomerAsset(g *gin.Context) {
	var request models.CustomerAsset
	g.Bind(&request)

	message, err := repository.DeleteCustomerAsset(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
