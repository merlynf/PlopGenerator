package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllHotelTypeDivision(g *gin.Context) {
	var request models.HotelTypeDivision
	g.Bind(&request)

	result, err := repository.GetAllHotelTypeDivision(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetHotelTypeDivision(g *gin.Context) {
	var request models.HotelTypeDivision
	g.Bind(&request)

	result, err := repository.GetHotelTypeDivision(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateHotelTypeDivision(g *gin.Context) {
	var request models.HotelTypeDivision
	g.Bind(&request)

	message, err := repository.CreateHotelTypeDivision(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateHotelTypeDivision(g *gin.Context) {
	var request models.HotelTypeDivision
	g.Bind(&request)

	message, err := repository.UpdateHotelTypeDivision(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteHotelTypeDivision(g *gin.Context) {
	var request models.HotelTypeDivision
	g.Bind(&request)

	message, err := repository.DeleteHotelTypeDivision(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
