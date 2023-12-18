package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllHotelRevenueSource(g *gin.Context) {
	var request models.HotelRevenueSource
	g.Bind(&request)

	result, err := repository.GetAllHotelRevenueSource(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetHotelRevenueSource(g *gin.Context) {
	var request models.HotelRevenueSource
	g.Bind(&request)

	result, err := repository.GetHotelRevenueSource(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateHotelRevenueSource(g *gin.Context) {
	var request models.HotelRevenueSource
	g.Bind(&request)

	message, err := repository.CreateHotelRevenueSource(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateHotelRevenueSource(g *gin.Context) {
	var request models.HotelRevenueSource
	g.Bind(&request)

	message, err := repository.UpdateHotelRevenueSource(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteHotelRevenueSource(g *gin.Context) {
	var request models.HotelRevenueSource
	g.Bind(&request)

	message, err := repository.DeleteHotelRevenueSource(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
