package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllHotelGroupUserHotel(g *gin.Context) {
	var request models.HotelGroupUserHotel
	g.Bind(&request)

	result, err := repository.GetAllHotelGroupUserHotel(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetHotelGroupUserHotel(g *gin.Context) {
	var request models.HotelGroupUserHotel
	g.Bind(&request)

	result, err := repository.GetHotelGroupUserHotel(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateHotelGroupUserHotel(g *gin.Context) {
	var request models.HotelGroupUserHotel
	g.Bind(&request)

	message, err := repository.CreateHotelGroupUserHotel(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateHotelGroupUserHotel(g *gin.Context) {
	var request models.HotelGroupUserHotel
	g.Bind(&request)

	message, err := repository.UpdateHotelGroupUserHotel(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteHotelGroupUserHotel(g *gin.Context) {
	var request models.HotelGroupUserHotel
	g.Bind(&request)

	message, err := repository.DeleteHotelGroupUserHotel(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
