package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllTaskDueOffsetDropTime(g *gin.Context) {
	var request models.TaskDueOffsetDropTime
	g.Bind(&request)

	result, err := repository.GetAllTaskDueOffsetDropTime(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetTaskDueOffsetDropTime(g *gin.Context) {
	var request models.TaskDueOffsetDropTime
	g.Bind(&request)

	result, err := repository.GetTaskDueOffsetDropTime(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateTaskDueOffsetDropTime(g *gin.Context) {
	var request models.TaskDueOffsetDropTime
	g.Bind(&request)

	message, err := repository.CreateTaskDueOffsetDropTime(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateTaskDueOffsetDropTime(g *gin.Context) {
	var request models.TaskDueOffsetDropTime
	g.Bind(&request)

	message, err := repository.UpdateTaskDueOffsetDropTime(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteTaskDueOffsetDropTime(g *gin.Context) {
	var request models.TaskDueOffsetDropTime
	g.Bind(&request)

	message, err := repository.DeleteTaskDueOffsetDropTime(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
