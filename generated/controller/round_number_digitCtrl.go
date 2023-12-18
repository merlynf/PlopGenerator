package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllRoundNumberDigit(g *gin.Context) {
	var request models.RoundNumberDigit
	g.Bind(&request)

	result, err := repository.GetAllRoundNumberDigit(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetRoundNumberDigit(g *gin.Context) {
	var request models.RoundNumberDigit
	g.Bind(&request)

	result, err := repository.GetRoundNumberDigit(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateRoundNumberDigit(g *gin.Context) {
	var request models.RoundNumberDigit
	g.Bind(&request)

	message, err := repository.CreateRoundNumberDigit(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateRoundNumberDigit(g *gin.Context) {
	var request models.RoundNumberDigit
	g.Bind(&request)

	message, err := repository.UpdateRoundNumberDigit(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteRoundNumberDigit(g *gin.Context) {
	var request models.RoundNumberDigit
	g.Bind(&request)

	message, err := repository.DeleteRoundNumberDigit(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
