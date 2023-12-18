package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllDateOfTheMonth(g *gin.Context) {
	var request models.DateOfTheMonth
	g.Bind(&request)

	result, err := repository.GetAllDateOfTheMonth(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetDateOfTheMonth(g *gin.Context) {
	var request models.DateOfTheMonth
	g.Bind(&request)

	result, err := repository.GetDateOfTheMonth(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateDateOfTheMonth(g *gin.Context) {
	var request models.DateOfTheMonth
	g.Bind(&request)

	message, err := repository.CreateDateOfTheMonth(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateDateOfTheMonth(g *gin.Context) {
	var request models.DateOfTheMonth
	g.Bind(&request)

	message, err := repository.UpdateDateOfTheMonth(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteDateOfTheMonth(g *gin.Context) {
	var request models.DateOfTheMonth
	g.Bind(&request)

	message, err := repository.DeleteDateOfTheMonth(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
