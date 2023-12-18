package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllCustomerAssignment(g *gin.Context) {
	var request models.CustomerAssignment
	g.Bind(&request)

	result, err := repository.GetAllCustomerAssignment(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetCustomerAssignment(g *gin.Context) {
	var request models.CustomerAssignment
	g.Bind(&request)

	result, err := repository.GetCustomerAssignment(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateCustomerAssignment(g *gin.Context) {
	var request models.CustomerAssignment
	g.Bind(&request)

	message, err := repository.CreateCustomerAssignment(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateCustomerAssignment(g *gin.Context) {
	var request models.CustomerAssignment
	g.Bind(&request)

	message, err := repository.UpdateCustomerAssignment(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteCustomerAssignment(g *gin.Context) {
	var request models.CustomerAssignment
	g.Bind(&request)

	message, err := repository.DeleteCustomerAssignment(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
