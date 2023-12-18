package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllHotelGroupHotel(g *gin.Context) {
	var request models.HotelGroupHotel
	g.Bind(&request)

	result, err := repository.GetAllHotelGroupHotel(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetHotelGroupHotel(g *gin.Context) {
	var request models.HotelGroupHotel
	g.Bind(&request)

	result, err := repository.GetHotelGroupHotel(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateHotelGroupHotel(g *gin.Context) {
	var request models.HotelGroupHotel
	g.Bind(&request)

	message, err := repository.CreateHotelGroupHotel(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateHotelGroupHotel(g *gin.Context) {
	var request models.HotelGroupHotel
	g.Bind(&request)

	message, err := repository.UpdateHotelGroupHotel(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteHotelGroupHotel(g *gin.Context) {
	var request models.HotelGroupHotel
	g.Bind(&request)

	message, err := repository.DeleteHotelGroupHotel(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
