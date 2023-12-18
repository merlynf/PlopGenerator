package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllCustomerContact(g *gin.Context) {
	var request models.CustomerContact
	g.Bind(&request)

	result, err := repository.GetAllCustomerContact(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetCustomerContact(g *gin.Context) {
	var request models.CustomerContact
	g.Bind(&request)

	result, err := repository.GetCustomerContact(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateCustomerContact(g *gin.Context) {
	var request models.CustomerContact
	g.Bind(&request)

	message, err := repository.CreateCustomerContact(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateCustomerContact(g *gin.Context) {
	var request models.CustomerContact
	g.Bind(&request)

	message, err := repository.UpdateCustomerContact(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteCustomerContact(g *gin.Context) {
	var request models.CustomerContact
	g.Bind(&request)

	message, err := repository.DeleteCustomerContact(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
