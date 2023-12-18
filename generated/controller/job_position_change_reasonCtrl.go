package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllJobPositionChangeReason(g *gin.Context) {
	var request models.JobPositionChangeReason
	g.Bind(&request)

	result, err := repository.GetAllJobPositionChangeReason(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetJobPositionChangeReason(g *gin.Context) {
	var request models.JobPositionChangeReason
	g.Bind(&request)

	result, err := repository.GetJobPositionChangeReason(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateJobPositionChangeReason(g *gin.Context) {
	var request models.JobPositionChangeReason
	g.Bind(&request)

	message, err := repository.CreateJobPositionChangeReason(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateJobPositionChangeReason(g *gin.Context) {
	var request models.JobPositionChangeReason
	g.Bind(&request)

	message, err := repository.UpdateJobPositionChangeReason(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteJobPositionChangeReason(g *gin.Context) {
	var request models.JobPositionChangeReason
	g.Bind(&request)

	message, err := repository.DeleteJobPositionChangeReason(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
