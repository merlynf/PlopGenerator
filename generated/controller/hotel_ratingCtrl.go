package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllHotelRating(g *gin.Context) {
	var request models.HotelRating
	g.Bind(&request)

	result, err := repository.GetAllHotelRating(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetHotelRating(g *gin.Context) {
	var request models.HotelRating
	g.Bind(&request)

	result, err := repository.GetHotelRating(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateHotelRating(g *gin.Context) {
	var request models.HotelRating
	g.Bind(&request)

	message, err := repository.CreateHotelRating(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateHotelRating(g *gin.Context) {
	var request models.HotelRating
	g.Bind(&request)

	message, err := repository.UpdateHotelRating(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteHotelRating(g *gin.Context) {
	var request models.HotelRating
	g.Bind(&request)

	message, err := repository.DeleteHotelRating(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
