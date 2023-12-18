package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllHotelGroupTypePage(g *gin.Context) {
	var request models.HotelGroupTypePage
	g.Bind(&request)

	result, err := repository.GetAllHotelGroupTypePage(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetHotelGroupTypePage(g *gin.Context) {
	var request models.HotelGroupTypePage
	g.Bind(&request)

	result, err := repository.GetHotelGroupTypePage(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateHotelGroupTypePage(g *gin.Context) {
	var request models.HotelGroupTypePage
	g.Bind(&request)

	message, err := repository.CreateHotelGroupTypePage(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateHotelGroupTypePage(g *gin.Context) {
	var request models.HotelGroupTypePage
	g.Bind(&request)

	message, err := repository.UpdateHotelGroupTypePage(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteHotelGroupTypePage(g *gin.Context) {
	var request models.HotelGroupTypePage
	g.Bind(&request)

	message, err := repository.DeleteHotelGroupTypePage(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
