package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllHotelJobTitle(g *gin.Context) {
	var request models.HotelJobTitle
	g.Bind(&request)

	result, err := repository.GetAllHotelJobTitle(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetHotelJobTitle(g *gin.Context) {
	var request models.HotelJobTitle
	g.Bind(&request)

	result, err := repository.GetHotelJobTitle(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateHotelJobTitle(g *gin.Context) {
	var request models.HotelJobTitle
	g.Bind(&request)

	message, err := repository.CreateHotelJobTitle(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateHotelJobTitle(g *gin.Context) {
	var request models.HotelJobTitle
	g.Bind(&request)

	message, err := repository.UpdateHotelJobTitle(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteHotelJobTitle(g *gin.Context) {
	var request models.HotelJobTitle
	g.Bind(&request)

	message, err := repository.DeleteHotelJobTitle(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
