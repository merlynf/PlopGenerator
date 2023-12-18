package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllHotelGroupJobTitle(g *gin.Context) {
	var request models.HotelGroupJobTitle
	g.Bind(&request)

	result, err := repository.GetAllHotelGroupJobTitle(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetHotelGroupJobTitle(g *gin.Context) {
	var request models.HotelGroupJobTitle
	g.Bind(&request)

	result, err := repository.GetHotelGroupJobTitle(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateHotelGroupJobTitle(g *gin.Context) {
	var request models.HotelGroupJobTitle
	g.Bind(&request)

	message, err := repository.CreateHotelGroupJobTitle(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateHotelGroupJobTitle(g *gin.Context) {
	var request models.HotelGroupJobTitle
	g.Bind(&request)

	message, err := repository.UpdateHotelGroupJobTitle(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteHotelGroupJobTitle(g *gin.Context) {
	var request models.HotelGroupJobTitle
	g.Bind(&request)

	message, err := repository.DeleteHotelGroupJobTitle(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
