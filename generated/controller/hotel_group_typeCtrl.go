package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllHotelGroupType(g *gin.Context) {
	var request models.HotelGroupType
	g.Bind(&request)

	result, err := repository.GetAllHotelGroupType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetHotelGroupType(g *gin.Context) {
	var request models.HotelGroupType
	g.Bind(&request)

	result, err := repository.GetHotelGroupType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateHotelGroupType(g *gin.Context) {
	var request models.HotelGroupType
	g.Bind(&request)

	message, err := repository.CreateHotelGroupType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateHotelGroupType(g *gin.Context) {
	var request models.HotelGroupType
	g.Bind(&request)

	message, err := repository.UpdateHotelGroupType(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteHotelGroupType(g *gin.Context) {
	var request models.HotelGroupType
	g.Bind(&request)

	message, err := repository.DeleteHotelGroupType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
