package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllHotelGroupTypeModule(g *gin.Context) {
	var request models.HotelGroupTypeModule
	g.Bind(&request)

	result, err := repository.GetAllHotelGroupTypeModule(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetHotelGroupTypeModule(g *gin.Context) {
	var request models.HotelGroupTypeModule
	g.Bind(&request)

	result, err := repository.GetHotelGroupTypeModule(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateHotelGroupTypeModule(g *gin.Context) {
	var request models.HotelGroupTypeModule
	g.Bind(&request)

	message, err := repository.CreateHotelGroupTypeModule(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateHotelGroupTypeModule(g *gin.Context) {
	var request models.HotelGroupTypeModule
	g.Bind(&request)

	message, err := repository.UpdateHotelGroupTypeModule(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteHotelGroupTypeModule(g *gin.Context) {
	var request models.HotelGroupTypeModule
	g.Bind(&request)

	message, err := repository.DeleteHotelGroupTypeModule(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
