package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllHotelTypeModule(g *gin.Context) {
	var request models.HotelTypeModule
	g.Bind(&request)

	result, err := repository.GetAllHotelTypeModule(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetHotelTypeModule(g *gin.Context) {
	var request models.HotelTypeModule
	g.Bind(&request)

	result, err := repository.GetHotelTypeModule(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateHotelTypeModule(g *gin.Context) {
	var request models.HotelTypeModule
	g.Bind(&request)

	message, err := repository.CreateHotelTypeModule(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateHotelTypeModule(g *gin.Context) {
	var request models.HotelTypeModule
	g.Bind(&request)

	message, err := repository.UpdateHotelTypeModule(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteHotelTypeModule(g *gin.Context) {
	var request models.HotelTypeModule
	g.Bind(&request)

	message, err := repository.DeleteHotelTypeModule(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
