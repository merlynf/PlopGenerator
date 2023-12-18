package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllRoundNumberFunction(g *gin.Context) {
	var request models.RoundNumberFunction
	g.Bind(&request)

	result, err := repository.GetAllRoundNumberFunction(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetRoundNumberFunction(g *gin.Context) {
	var request models.RoundNumberFunction
	g.Bind(&request)

	result, err := repository.GetRoundNumberFunction(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateRoundNumberFunction(g *gin.Context) {
	var request models.RoundNumberFunction
	g.Bind(&request)

	message, err := repository.CreateRoundNumberFunction(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateRoundNumberFunction(g *gin.Context) {
	var request models.RoundNumberFunction
	g.Bind(&request)

	message, err := repository.UpdateRoundNumberFunction(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteRoundNumberFunction(g *gin.Context) {
	var request models.RoundNumberFunction
	g.Bind(&request)

	message, err := repository.DeleteRoundNumberFunction(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
