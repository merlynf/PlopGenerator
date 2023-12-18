package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllHotelUser(g *gin.Context) {
	var request models.HotelUser
	g.Bind(&request)

	result, err := repository.GetAllHotelUser(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetHotelUser(g *gin.Context) {
	var request models.HotelUser
	g.Bind(&request)

	result, err := repository.GetHotelUser(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateHotelUser(g *gin.Context) {
	var request models.HotelUser
	g.Bind(&request)

	message, err := repository.CreateHotelUser(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateHotelUser(g *gin.Context) {
	var request models.HotelUser
	g.Bind(&request)

	message, err := repository.UpdateHotelUser(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteHotelUser(g *gin.Context) {
	var request models.HotelUser
	g.Bind(&request)

	message, err := repository.DeleteHotelUser(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
