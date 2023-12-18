package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllHotelHotelRevenueSource(g *gin.Context) {
	var request models.HotelHotelRevenueSource
	g.Bind(&request)

	result, err := repository.GetAllHotelHotelRevenueSource(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetHotelHotelRevenueSource(g *gin.Context) {
	var request models.HotelHotelRevenueSource
	g.Bind(&request)

	result, err := repository.GetHotelHotelRevenueSource(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateHotelHotelRevenueSource(g *gin.Context) {
	var request models.HotelHotelRevenueSource
	g.Bind(&request)

	message, err := repository.CreateHotelHotelRevenueSource(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateHotelHotelRevenueSource(g *gin.Context) {
	var request models.HotelHotelRevenueSource
	g.Bind(&request)

	message, err := repository.UpdateHotelHotelRevenueSource(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteHotelHotelRevenueSource(g *gin.Context) {
	var request models.HotelHotelRevenueSource
	g.Bind(&request)

	message, err := repository.DeleteHotelHotelRevenueSource(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
