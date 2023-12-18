package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllHotelPage(g *gin.Context) {
	var request models.HotelPage
	g.Bind(&request)

	result, err := repository.GetAllHotelPage(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetHotelPage(g *gin.Context) {
	var request models.HotelPage
	g.Bind(&request)

	result, err := repository.GetHotelPage(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateHotelPage(g *gin.Context) {
	var request models.HotelPage
	g.Bind(&request)

	message, err := repository.CreateHotelPage(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateHotelPage(g *gin.Context) {
	var request models.HotelPage
	g.Bind(&request)

	message, err := repository.UpdateHotelPage(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteHotelPage(g *gin.Context) {
	var request models.HotelPage
	g.Bind(&request)

	message, err := repository.DeleteHotelPage(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
