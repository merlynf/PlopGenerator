package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllHotelTypePage(g *gin.Context) {
	var request models.HotelTypePage
	g.Bind(&request)

	result, err := repository.GetAllHotelTypePage(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetHotelTypePage(g *gin.Context) {
	var request models.HotelTypePage
	g.Bind(&request)

	result, err := repository.GetHotelTypePage(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateHotelTypePage(g *gin.Context) {
	var request models.HotelTypePage
	g.Bind(&request)

	message, err := repository.CreateHotelTypePage(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateHotelTypePage(g *gin.Context) {
	var request models.HotelTypePage
	g.Bind(&request)

	message, err := repository.UpdateHotelTypePage(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteHotelTypePage(g *gin.Context) {
	var request models.HotelTypePage
	g.Bind(&request)

	message, err := repository.DeleteHotelTypePage(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
