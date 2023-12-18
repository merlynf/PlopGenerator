package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllEmploymentType(g *gin.Context) {
	var request models.EmploymentType
	g.Bind(&request)

	result, err := repository.GetAllEmploymentType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetEmploymentType(g *gin.Context) {
	var request models.EmploymentType
	g.Bind(&request)

	result, err := repository.GetEmploymentType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateEmploymentType(g *gin.Context) {
	var request models.EmploymentType
	g.Bind(&request)

	message, err := repository.CreateEmploymentType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateEmploymentType(g *gin.Context) {
	var request models.EmploymentType
	g.Bind(&request)

	message, err := repository.UpdateEmploymentType(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteEmploymentType(g *gin.Context) {
	var request models.EmploymentType
	g.Bind(&request)

	message, err := repository.DeleteEmploymentType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
