package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllHotelSalesperson(g *gin.Context) {
	var request models.HotelSalesperson
	g.Bind(&request)

	result, err := repository.GetAllHotelSalesperson(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetHotelSalesperson(g *gin.Context) {
	var request models.HotelSalesperson
	g.Bind(&request)

	result, err := repository.GetHotelSalesperson(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateHotelSalesperson(g *gin.Context) {
	var request models.HotelSalesperson
	g.Bind(&request)

	message, err := repository.CreateHotelSalesperson(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateHotelSalesperson(g *gin.Context) {
	var request models.HotelSalesperson
	g.Bind(&request)

	message, err := repository.UpdateHotelSalesperson(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteHotelSalesperson(g *gin.Context) {
	var request models.HotelSalesperson
	g.Bind(&request)

	message, err := repository.DeleteHotelSalesperson(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
