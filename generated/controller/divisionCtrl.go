package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllDivision(g *gin.Context) {
	var request models.Division
	g.Bind(&request)

	result, err := repository.GetAllDivision(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetDivision(g *gin.Context) {
	var request models.Division
	g.Bind(&request)

	result, err := repository.GetDivision(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateDivision(g *gin.Context) {
	var request models.Division
	g.Bind(&request)

	message, err := repository.CreateDivision(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateDivision(g *gin.Context) {
	var request models.Division
	g.Bind(&request)

	message, err := repository.UpdateDivision(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteDivision(g *gin.Context) {
	var request models.Division
	g.Bind(&request)

	message, err := repository.DeleteDivision(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
