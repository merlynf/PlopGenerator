package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllEmployeeAsset(g *gin.Context) {
	var request models.EmployeeAsset
	g.Bind(&request)

	result, err := repository.GetAllEmployeeAsset(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetEmployeeAsset(g *gin.Context) {
	var request models.EmployeeAsset
	g.Bind(&request)

	result, err := repository.GetEmployeeAsset(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateEmployeeAsset(g *gin.Context) {
	var request models.EmployeeAsset
	g.Bind(&request)

	message, err := repository.CreateEmployeeAsset(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateEmployeeAsset(g *gin.Context) {
	var request models.EmployeeAsset
	g.Bind(&request)

	message, err := repository.UpdateEmployeeAsset(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteEmployeeAsset(g *gin.Context) {
	var request models.EmployeeAsset
	g.Bind(&request)

	message, err := repository.DeleteEmployeeAsset(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
