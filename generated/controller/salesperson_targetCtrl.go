package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllSalespersonTarget(g *gin.Context) {
	var request models.SalespersonTarget
	g.Bind(&request)

	result, err := repository.GetAllSalespersonTarget(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetSalespersonTarget(g *gin.Context) {
	var request models.SalespersonTarget
	g.Bind(&request)

	result, err := repository.GetSalespersonTarget(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateSalespersonTarget(g *gin.Context) {
	var request models.SalespersonTarget
	g.Bind(&request)

	message, err := repository.CreateSalespersonTarget(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateSalespersonTarget(g *gin.Context) {
	var request models.SalespersonTarget
	g.Bind(&request)

	message, err := repository.UpdateSalespersonTarget(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteSalespersonTarget(g *gin.Context) {
	var request models.SalespersonTarget
	g.Bind(&request)

	message, err := repository.DeleteSalespersonTarget(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
