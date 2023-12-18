package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllSalesProspectContact(g *gin.Context) {
	var request models.SalesProspectContact
	g.Bind(&request)

	result, err := repository.GetAllSalesProspectContact(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetSalesProspectContact(g *gin.Context) {
	var request models.SalesProspectContact
	g.Bind(&request)

	result, err := repository.GetSalesProspectContact(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateSalesProspectContact(g *gin.Context) {
	var request models.SalesProspectContact
	g.Bind(&request)

	message, err := repository.CreateSalesProspectContact(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateSalesProspectContact(g *gin.Context) {
	var request models.SalesProspectContact
	g.Bind(&request)

	message, err := repository.UpdateSalesProspectContact(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteSalesProspectContact(g *gin.Context) {
	var request models.SalesProspectContact
	g.Bind(&request)

	message, err := repository.DeleteSalesProspectContact(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
