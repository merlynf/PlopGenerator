package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllHotelTypeHotelRevenueSource(g *gin.Context) {
	var request models.HotelTypeHotelRevenueSource
	g.Bind(&request)

	result, err := repository.GetAllHotelTypeHotelRevenueSource(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetHotelTypeHotelRevenueSource(g *gin.Context) {
	var request models.HotelTypeHotelRevenueSource
	g.Bind(&request)

	result, err := repository.GetHotelTypeHotelRevenueSource(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateHotelTypeHotelRevenueSource(g *gin.Context) {
	var request models.HotelTypeHotelRevenueSource
	g.Bind(&request)

	message, err := repository.CreateHotelTypeHotelRevenueSource(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateHotelTypeHotelRevenueSource(g *gin.Context) {
	var request models.HotelTypeHotelRevenueSource
	g.Bind(&request)

	message, err := repository.UpdateHotelTypeHotelRevenueSource(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteHotelTypeHotelRevenueSource(g *gin.Context) {
	var request models.HotelTypeHotelRevenueSource
	g.Bind(&request)

	message, err := repository.DeleteHotelTypeHotelRevenueSource(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
