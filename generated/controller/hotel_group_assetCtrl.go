package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllHotelGroupAsset(g *gin.Context) {
	var request models.HotelGroupAsset
	g.Bind(&request)

	result, err := repository.GetAllHotelGroupAsset(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetHotelGroupAsset(g *gin.Context) {
	var request models.HotelGroupAsset
	g.Bind(&request)

	result, err := repository.GetHotelGroupAsset(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateHotelGroupAsset(g *gin.Context) {
	var request models.HotelGroupAsset
	g.Bind(&request)

	message, err := repository.CreateHotelGroupAsset(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateHotelGroupAsset(g *gin.Context) {
	var request models.HotelGroupAsset
	g.Bind(&request)

	message, err := repository.UpdateHotelGroupAsset(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteHotelGroupAsset(g *gin.Context) {
	var request models.HotelGroupAsset
	g.Bind(&request)

	message, err := repository.DeleteHotelGroupAsset(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
