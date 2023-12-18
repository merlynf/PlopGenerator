package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllSystemTaskDueOffsetDropTime(g *gin.Context) {
	var request models.SystemTaskDueOffsetDropTime
	g.Bind(&request)

	result, err := repository.GetAllSystemTaskDueOffsetDropTime(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetSystemTaskDueOffsetDropTime(g *gin.Context) {
	var request models.SystemTaskDueOffsetDropTime
	g.Bind(&request)

	result, err := repository.GetSystemTaskDueOffsetDropTime(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateSystemTaskDueOffsetDropTime(g *gin.Context) {
	var request models.SystemTaskDueOffsetDropTime
	g.Bind(&request)

	message, err := repository.CreateSystemTaskDueOffsetDropTime(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateSystemTaskDueOffsetDropTime(g *gin.Context) {
	var request models.SystemTaskDueOffsetDropTime
	g.Bind(&request)

	message, err := repository.UpdateSystemTaskDueOffsetDropTime(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteSystemTaskDueOffsetDropTime(g *gin.Context) {
	var request models.SystemTaskDueOffsetDropTime
	g.Bind(&request)

	message, err := repository.DeleteSystemTaskDueOffsetDropTime(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
