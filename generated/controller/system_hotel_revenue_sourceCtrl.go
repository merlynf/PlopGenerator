package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllSystemHotelRevenueSource(g *gin.Context) {
	var request models.SystemHotelRevenueSource
	g.Bind(&request)

	result, err := repository.GetAllSystemHotelRevenueSource(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetSystemHotelRevenueSource(g *gin.Context) {
	var request models.SystemHotelRevenueSource
	g.Bind(&request)

	result, err := repository.GetSystemHotelRevenueSource(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateSystemHotelRevenueSource(g *gin.Context) {
	var request models.SystemHotelRevenueSource
	g.Bind(&request)

	message, err := repository.CreateSystemHotelRevenueSource(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateSystemHotelRevenueSource(g *gin.Context) {
	var request models.SystemHotelRevenueSource
	g.Bind(&request)

	message, err := repository.UpdateSystemHotelRevenueSource(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteSystemHotelRevenueSource(g *gin.Context) {
	var request models.SystemHotelRevenueSource
	g.Bind(&request)

	message, err := repository.DeleteSystemHotelRevenueSource(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
