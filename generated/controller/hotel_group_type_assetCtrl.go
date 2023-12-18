package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllHotelGroupTypeAsset(g *gin.Context) {
	var request models.HotelGroupTypeAsset
	g.Bind(&request)

	result, err := repository.GetAllHotelGroupTypeAsset(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetHotelGroupTypeAsset(g *gin.Context) {
	var request models.HotelGroupTypeAsset
	g.Bind(&request)

	result, err := repository.GetHotelGroupTypeAsset(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateHotelGroupTypeAsset(g *gin.Context) {
	var request models.HotelGroupTypeAsset
	g.Bind(&request)

	message, err := repository.CreateHotelGroupTypeAsset(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateHotelGroupTypeAsset(g *gin.Context) {
	var request models.HotelGroupTypeAsset
	g.Bind(&request)

	message, err := repository.UpdateHotelGroupTypeAsset(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteHotelGroupTypeAsset(g *gin.Context) {
	var request models.HotelGroupTypeAsset
	g.Bind(&request)

	message, err := repository.DeleteHotelGroupTypeAsset(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
