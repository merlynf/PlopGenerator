package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllHotelTypeJobTitle(g *gin.Context) {
	var request models.HotelTypeJobTitle
	g.Bind(&request)

	result, err := repository.GetAllHotelTypeJobTitle(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetHotelTypeJobTitle(g *gin.Context) {
	var request models.HotelTypeJobTitle
	g.Bind(&request)

	result, err := repository.GetHotelTypeJobTitle(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateHotelTypeJobTitle(g *gin.Context) {
	var request models.HotelTypeJobTitle
	g.Bind(&request)

	message, err := repository.CreateHotelTypeJobTitle(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateHotelTypeJobTitle(g *gin.Context) {
	var request models.HotelTypeJobTitle
	g.Bind(&request)

	message, err := repository.UpdateHotelTypeJobTitle(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteHotelTypeJobTitle(g *gin.Context) {
	var request models.HotelTypeJobTitle
	g.Bind(&request)

	message, err := repository.DeleteHotelTypeJobTitle(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
