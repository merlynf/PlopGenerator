package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllEmployeeAdditionalRole(g *gin.Context) {
	var request models.EmployeeAdditionalRole
	g.Bind(&request)

	result, err := repository.GetAllEmployeeAdditionalRole(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetEmployeeAdditionalRole(g *gin.Context) {
	var request models.EmployeeAdditionalRole
	g.Bind(&request)

	result, err := repository.GetEmployeeAdditionalRole(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateEmployeeAdditionalRole(g *gin.Context) {
	var request models.EmployeeAdditionalRole
	g.Bind(&request)

	message, err := repository.CreateEmployeeAdditionalRole(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateEmployeeAdditionalRole(g *gin.Context) {
	var request models.EmployeeAdditionalRole
	g.Bind(&request)

	message, err := repository.UpdateEmployeeAdditionalRole(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteEmployeeAdditionalRole(g *gin.Context) {
	var request models.EmployeeAdditionalRole
	g.Bind(&request)

	message, err := repository.DeleteEmployeeAdditionalRole(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
