package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllRatingType(g *gin.Context) {
	var request models.RatingType
	g.Bind(&request)

	result, err := repository.GetAllRatingType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetRatingType(g *gin.Context) {
	var request models.RatingType
	g.Bind(&request)

	result, err := repository.GetRatingType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateRatingType(g *gin.Context) {
	var request models.RatingType
	g.Bind(&request)

	message, err := repository.CreateRatingType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateRatingType(g *gin.Context) {
	var request models.RatingType
	g.Bind(&request)

	message, err := repository.UpdateRatingType(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteRatingType(g *gin.Context) {
	var request models.RatingType
	g.Bind(&request)

	message, err := repository.DeleteRatingType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
