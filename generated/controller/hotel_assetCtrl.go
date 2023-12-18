package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllHotelAsset(g *gin.Context) {
	var request models.HotelAsset
	g.Bind(&request)

	result, err := repository.GetAllHotelAsset(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetHotelAsset(g *gin.Context) {
	var request models.HotelAsset
	g.Bind(&request)

	result, err := repository.GetHotelAsset(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateHotelAsset(g *gin.Context) {
	var request models.HotelAsset
	g.Bind(&request)

	message, err := repository.CreateHotelAsset(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateHotelAsset(g *gin.Context) {
	var request models.HotelAsset
	g.Bind(&request)

	message, err := repository.UpdateHotelAsset(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteHotelAsset(g *gin.Context) {
	var request models.HotelAsset
	g.Bind(&request)

	message, err := repository.DeleteHotelAsset(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
