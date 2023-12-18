package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllHotelGroupUser(g *gin.Context) {
	var request models.HotelGroupUser
	g.Bind(&request)

	result, err := repository.GetAllHotelGroupUser(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetHotelGroupUser(g *gin.Context) {
	var request models.HotelGroupUser
	g.Bind(&request)

	result, err := repository.GetHotelGroupUser(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateHotelGroupUser(g *gin.Context) {
	var request models.HotelGroupUser
	g.Bind(&request)

	message, err := repository.CreateHotelGroupUser(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateHotelGroupUser(g *gin.Context) {
	var request models.HotelGroupUser
	g.Bind(&request)

	message, err := repository.UpdateHotelGroupUser(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteHotelGroupUser(g *gin.Context) {
	var request models.HotelGroupUser
	g.Bind(&request)

	message, err := repository.DeleteHotelGroupUser(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
