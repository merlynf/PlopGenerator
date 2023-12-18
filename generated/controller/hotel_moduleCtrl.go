package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllHotelModule(g *gin.Context) {
	var request models.HotelModule
	g.Bind(&request)

	result, err := repository.GetAllHotelModule(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetHotelModule(g *gin.Context) {
	var request models.HotelModule
	g.Bind(&request)

	result, err := repository.GetHotelModule(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateHotelModule(g *gin.Context) {
	var request models.HotelModule
	g.Bind(&request)

	message, err := repository.CreateHotelModule(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateHotelModule(g *gin.Context) {
	var request models.HotelModule
	g.Bind(&request)

	message, err := repository.UpdateHotelModule(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteHotelModule(g *gin.Context) {
	var request models.HotelModule
	g.Bind(&request)

	message, err := repository.DeleteHotelModule(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
