package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllHotelGroupContact(g *gin.Context) {
	var request models.HotelGroupContact
	g.Bind(&request)

	result, err := repository.GetAllHotelGroupContact(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetHotelGroupContact(g *gin.Context) {
	var request models.HotelGroupContact
	g.Bind(&request)

	result, err := repository.GetHotelGroupContact(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateHotelGroupContact(g *gin.Context) {
	var request models.HotelGroupContact
	g.Bind(&request)

	message, err := repository.CreateHotelGroupContact(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateHotelGroupContact(g *gin.Context) {
	var request models.HotelGroupContact
	g.Bind(&request)

	message, err := repository.UpdateHotelGroupContact(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteHotelGroupContact(g *gin.Context) {
	var request models.HotelGroupContact
	g.Bind(&request)

	message, err := repository.DeleteHotelGroupContact(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
