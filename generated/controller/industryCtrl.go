package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllIndustry(g *gin.Context) {
	var request models.Industry
	g.Bind(&request)

	result, err := repository.GetAllIndustry(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetIndustry(g *gin.Context) {
	var request models.Industry
	g.Bind(&request)

	result, err := repository.GetIndustry(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateIndustry(g *gin.Context) {
	var request models.Industry
	g.Bind(&request)

	message, err := repository.CreateIndustry(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateIndustry(g *gin.Context) {
	var request models.Industry
	g.Bind(&request)

	message, err := repository.UpdateIndustry(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteIndustry(g *gin.Context) {
	var request models.Industry
	g.Bind(&request)

	message, err := repository.DeleteIndustry(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
