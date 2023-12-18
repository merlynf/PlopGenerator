package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllSpecialRequest(g *gin.Context) {
	var request models.SpecialRequest
	g.Bind(&request)

	result, err := repository.GetAllSpecialRequest(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetSpecialRequest(g *gin.Context) {
	var request models.SpecialRequest
	g.Bind(&request)

	result, err := repository.GetSpecialRequest(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateSpecialRequest(g *gin.Context) {
	var request models.SpecialRequest
	g.Bind(&request)

	message, err := repository.CreateSpecialRequest(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateSpecialRequest(g *gin.Context) {
	var request models.SpecialRequest
	g.Bind(&request)

	message, err := repository.UpdateSpecialRequest(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteSpecialRequest(g *gin.Context) {
	var request models.SpecialRequest
	g.Bind(&request)

	message, err := repository.DeleteSpecialRequest(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
