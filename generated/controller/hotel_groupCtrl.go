package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllHotelGroup(g *gin.Context) {
	var request models.HotelGroup
	g.Bind(&request)

	result, err := repository.GetAllHotelGroup(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetHotelGroup(g *gin.Context) {
	var request models.HotelGroup
	g.Bind(&request)

	result, err := repository.GetHotelGroup(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateHotelGroup(g *gin.Context) {
	var request models.HotelGroup
	g.Bind(&request)

	message, err := repository.CreateHotelGroup(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateHotelGroup(g *gin.Context) {
	var request models.HotelGroup
	g.Bind(&request)

	message, err := repository.UpdateHotelGroup(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteHotelGroup(g *gin.Context) {
	var request models.HotelGroup
	g.Bind(&request)

	message, err := repository.DeleteHotelGroup(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
