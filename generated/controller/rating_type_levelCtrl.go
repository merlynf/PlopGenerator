package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllRatingTypeLevel(g *gin.Context) {
	var request models.RatingTypeLevel
	g.Bind(&request)

	result, err := repository.GetAllRatingTypeLevel(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetRatingTypeLevel(g *gin.Context) {
	var request models.RatingTypeLevel
	g.Bind(&request)

	result, err := repository.GetRatingTypeLevel(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateRatingTypeLevel(g *gin.Context) {
	var request models.RatingTypeLevel
	g.Bind(&request)

	message, err := repository.CreateRatingTypeLevel(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateRatingTypeLevel(g *gin.Context) {
	var request models.RatingTypeLevel
	g.Bind(&request)

	message, err := repository.UpdateRatingTypeLevel(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteRatingTypeLevel(g *gin.Context) {
	var request models.RatingTypeLevel
	g.Bind(&request)

	message, err := repository.DeleteRatingTypeLevel(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
