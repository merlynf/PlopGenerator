package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllHotelGroupEmployee(g *gin.Context) {
	var request models.HotelGroupEmployee
	g.Bind(&request)

	result, err := repository.GetAllHotelGroupEmployee(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetHotelGroupEmployee(g *gin.Context) {
	var request models.HotelGroupEmployee
	g.Bind(&request)

	result, err := repository.GetHotelGroupEmployee(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateHotelGroupEmployee(g *gin.Context) {
	var request models.HotelGroupEmployee
	g.Bind(&request)

	message, err := repository.CreateHotelGroupEmployee(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateHotelGroupEmployee(g *gin.Context) {
	var request models.HotelGroupEmployee
	g.Bind(&request)

	message, err := repository.UpdateHotelGroupEmployee(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteHotelGroupEmployee(g *gin.Context) {
	var request models.HotelGroupEmployee
	g.Bind(&request)

	message, err := repository.DeleteHotelGroupEmployee(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
