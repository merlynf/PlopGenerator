package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllTaskStatus(g *gin.Context) {
	var request models.TaskStatus
	g.Bind(&request)

	result, err := repository.GetAllTaskStatus(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetTaskStatus(g *gin.Context) {
	var request models.TaskStatus
	g.Bind(&request)

	result, err := repository.GetTaskStatus(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateTaskStatus(g *gin.Context) {
	var request models.TaskStatus
	g.Bind(&request)

	message, err := repository.CreateTaskStatus(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateTaskStatus(g *gin.Context) {
	var request models.TaskStatus
	g.Bind(&request)

	message, err := repository.UpdateTaskStatus(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteTaskStatus(g *gin.Context) {
	var request models.TaskStatus
	g.Bind(&request)

	message, err := repository.DeleteTaskStatus(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
