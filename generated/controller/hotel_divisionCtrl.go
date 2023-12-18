package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllHotelDivision(g *gin.Context) {
	var request models.HotelDivision
	g.Bind(&request)

	result, err := repository.GetAllHotelDivision(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetHotelDivision(g *gin.Context) {
	var request models.HotelDivision
	g.Bind(&request)

	result, err := repository.GetHotelDivision(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateHotelDivision(g *gin.Context) {
	var request models.HotelDivision
	g.Bind(&request)

	message, err := repository.CreateHotelDivision(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateHotelDivision(g *gin.Context) {
	var request models.HotelDivision
	g.Bind(&request)

	message, err := repository.UpdateHotelDivision(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteHotelDivision(g *gin.Context) {
	var request models.HotelDivision
	g.Bind(&request)

	message, err := repository.DeleteHotelDivision(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
