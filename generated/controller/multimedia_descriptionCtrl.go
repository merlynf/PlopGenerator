package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllMultimediaDescription(g *gin.Context) {
	var request models.MultimediaDescription
	g.Bind(&request)

	result, err := repository.GetAllMultimediaDescription(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetMultimediaDescription(g *gin.Context) {
	var request models.MultimediaDescription
	g.Bind(&request)

	result, err := repository.GetMultimediaDescription(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateMultimediaDescription(g *gin.Context) {
	var request models.MultimediaDescription
	g.Bind(&request)

	message, err := repository.CreateMultimediaDescription(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateMultimediaDescription(g *gin.Context) {
	var request models.MultimediaDescription
	g.Bind(&request)

	message, err := repository.UpdateMultimediaDescription(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteMultimediaDescription(g *gin.Context) {
	var request models.MultimediaDescription
	g.Bind(&request)

	message, err := repository.DeleteMultimediaDescription(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
