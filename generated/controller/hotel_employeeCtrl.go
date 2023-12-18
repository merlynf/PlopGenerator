package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllHotelEmployee(g *gin.Context) {
	var request models.HotelEmployee
	g.Bind(&request)

	result, err := repository.GetAllHotelEmployee(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetHotelEmployee(g *gin.Context) {
	var request models.HotelEmployee
	g.Bind(&request)

	result, err := repository.GetHotelEmployee(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateHotelEmployee(g *gin.Context) {
	var request models.HotelEmployee
	g.Bind(&request)

	message, err := repository.CreateHotelEmployee(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateHotelEmployee(g *gin.Context) {
	var request models.HotelEmployee
	g.Bind(&request)

	message, err := repository.UpdateHotelEmployee(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteHotelEmployee(g *gin.Context) {
	var request models.HotelEmployee
	g.Bind(&request)

	message, err := repository.DeleteHotelEmployee(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
