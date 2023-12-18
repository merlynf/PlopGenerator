package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllEmployeeHistory(g *gin.Context) {
	var request models.EmployeeHistory
	g.Bind(&request)

	result, err := repository.GetAllEmployeeHistory(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetEmployeeHistory(g *gin.Context) {
	var request models.EmployeeHistory
	g.Bind(&request)

	result, err := repository.GetEmployeeHistory(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateEmployeeHistory(g *gin.Context) {
	var request models.EmployeeHistory
	g.Bind(&request)

	message, err := repository.CreateEmployeeHistory(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateEmployeeHistory(g *gin.Context) {
	var request models.EmployeeHistory
	g.Bind(&request)

	message, err := repository.UpdateEmployeeHistory(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteEmployeeHistory(g *gin.Context) {
	var request models.EmployeeHistory
	g.Bind(&request)

	message, err := repository.DeleteEmployeeHistory(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
