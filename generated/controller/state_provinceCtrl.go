package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllStateProvince(g *gin.Context) {
	var request models.StateProvince
	g.Bind(&request)

	result, err := repository.GetAllStateProvince(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetStateProvince(g *gin.Context) {
	var request models.StateProvince
	g.Bind(&request)

	result, err := repository.GetStateProvince(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateStateProvince(g *gin.Context) {
	var request models.StateProvince
	g.Bind(&request)

	message, err := repository.CreateStateProvince(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateStateProvince(g *gin.Context) {
	var request models.StateProvince
	g.Bind(&request)

	message, err := repository.UpdateStateProvince(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteStateProvince(g *gin.Context) {
	var request models.StateProvince
	g.Bind(&request)

	message, err := repository.DeleteStateProvince(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
