package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllPictureCategory(g *gin.Context) {
	var request models.PictureCategory
	g.Bind(&request)

	result, err := repository.GetAllPictureCategory(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetPictureCategory(g *gin.Context) {
	var request models.PictureCategory
	g.Bind(&request)

	result, err := repository.GetPictureCategory(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreatePictureCategory(g *gin.Context) {
	var request models.PictureCategory
	g.Bind(&request)

	message, err := repository.CreatePictureCategory(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdatePictureCategory(g *gin.Context) {
	var request models.PictureCategory
	g.Bind(&request)

	message, err := repository.UpdatePictureCategory(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeletePictureCategory(g *gin.Context) {
	var request models.PictureCategory
	g.Bind(&request)

	message, err := repository.DeletePictureCategory(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
