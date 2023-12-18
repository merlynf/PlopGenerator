package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllHotelContact(g *gin.Context) {
	var request models.HotelContact
	g.Bind(&request)

	result, err := repository.GetAllHotelContact(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetHotelContact(g *gin.Context) {
	var request models.HotelContact
	g.Bind(&request)

	result, err := repository.GetHotelContact(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateHotelContact(g *gin.Context) {
	var request models.HotelContact
	g.Bind(&request)

	message, err := repository.CreateHotelContact(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateHotelContact(g *gin.Context) {
	var request models.HotelContact
	g.Bind(&request)

	message, err := repository.UpdateHotelContact(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteHotelContact(g *gin.Context) {
	var request models.HotelContact
	g.Bind(&request)

	message, err := repository.DeleteHotelContact(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
