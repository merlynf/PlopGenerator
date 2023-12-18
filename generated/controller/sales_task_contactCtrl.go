package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllSalesTaskContact(g *gin.Context) {
	var request models.SalesTaskContact
	g.Bind(&request)

	result, err := repository.GetAllSalesTaskContact(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetSalesTaskContact(g *gin.Context) {
	var request models.SalesTaskContact
	g.Bind(&request)

	result, err := repository.GetSalesTaskContact(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateSalesTaskContact(g *gin.Context) {
	var request models.SalesTaskContact
	g.Bind(&request)

	message, err := repository.CreateSalesTaskContact(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateSalesTaskContact(g *gin.Context) {
	var request models.SalesTaskContact
	g.Bind(&request)

	message, err := repository.UpdateSalesTaskContact(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteSalesTaskContact(g *gin.Context) {
	var request models.SalesTaskContact
	g.Bind(&request)

	message, err := repository.DeleteSalesTaskContact(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
