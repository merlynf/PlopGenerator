package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllHotelUserType(g *gin.Context) {
	var request models.HotelUserType
	g.Bind(&request)

	result, err := repository.GetAllHotelUserType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetHotelUserType(g *gin.Context) {
	var request models.HotelUserType
	g.Bind(&request)

	result, err := repository.GetHotelUserType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateHotelUserType(g *gin.Context) {
	var request models.HotelUserType
	g.Bind(&request)

	message, err := repository.CreateHotelUserType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateHotelUserType(g *gin.Context) {
	var request models.HotelUserType
	g.Bind(&request)

	message, err := repository.UpdateHotelUserType(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteHotelUserType(g *gin.Context) {
	var request models.HotelUserType
	g.Bind(&request)

	message, err := repository.DeleteHotelUserType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
