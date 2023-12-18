package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllHotelTypeAsset(g *gin.Context) {
	var request models.HotelTypeAsset
	g.Bind(&request)

	result, err := repository.GetAllHotelTypeAsset(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetHotelTypeAsset(g *gin.Context) {
	var request models.HotelTypeAsset
	g.Bind(&request)

	result, err := repository.GetHotelTypeAsset(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateHotelTypeAsset(g *gin.Context) {
	var request models.HotelTypeAsset
	g.Bind(&request)

	message, err := repository.CreateHotelTypeAsset(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateHotelTypeAsset(g *gin.Context) {
	var request models.HotelTypeAsset
	g.Bind(&request)

	message, err := repository.UpdateHotelTypeAsset(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteHotelTypeAsset(g *gin.Context) {
	var request models.HotelTypeAsset
	g.Bind(&request)

	message, err := repository.DeleteHotelTypeAsset(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
