package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllHotelGroupTypeDivision(g *gin.Context) {
	var request models.HotelGroupTypeDivision
	g.Bind(&request)

	result, err := repository.GetAllHotelGroupTypeDivision(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetHotelGroupTypeDivision(g *gin.Context) {
	var request models.HotelGroupTypeDivision
	g.Bind(&request)

	result, err := repository.GetHotelGroupTypeDivision(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateHotelGroupTypeDivision(g *gin.Context) {
	var request models.HotelGroupTypeDivision
	g.Bind(&request)

	message, err := repository.CreateHotelGroupTypeDivision(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateHotelGroupTypeDivision(g *gin.Context) {
	var request models.HotelGroupTypeDivision
	g.Bind(&request)

	message, err := repository.UpdateHotelGroupTypeDivision(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteHotelGroupTypeDivision(g *gin.Context) {
	var request models.HotelGroupTypeDivision
	g.Bind(&request)

	message, err := repository.DeleteHotelGroupTypeDivision(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
