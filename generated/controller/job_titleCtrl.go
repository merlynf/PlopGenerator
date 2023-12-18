package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllJobTitle(g *gin.Context) {
	var request models.JobTitle
	g.Bind(&request)

	result, err := repository.GetAllJobTitle(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetJobTitle(g *gin.Context) {
	var request models.JobTitle
	g.Bind(&request)

	result, err := repository.GetJobTitle(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateJobTitle(g *gin.Context) {
	var request models.JobTitle
	g.Bind(&request)

	message, err := repository.CreateJobTitle(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateJobTitle(g *gin.Context) {
	var request models.JobTitle
	g.Bind(&request)

	message, err := repository.UpdateJobTitle(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteJobTitle(g *gin.Context) {
	var request models.JobTitle
	g.Bind(&request)

	message, err := repository.DeleteJobTitle(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
