package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllHotelGroupTypeUserType(g *gin.Context) {
	var request models.HotelGroupTypeUserType
	g.Bind(&request)

	result, err := repository.GetAllHotelGroupTypeUserType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetHotelGroupTypeUserType(g *gin.Context) {
	var request models.HotelGroupTypeUserType
	g.Bind(&request)

	result, err := repository.GetHotelGroupTypeUserType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateHotelGroupTypeUserType(g *gin.Context) {
	var request models.HotelGroupTypeUserType
	g.Bind(&request)

	message, err := repository.CreateHotelGroupTypeUserType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateHotelGroupTypeUserType(g *gin.Context) {
	var request models.HotelGroupTypeUserType
	g.Bind(&request)

	message, err := repository.UpdateHotelGroupTypeUserType(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteHotelGroupTypeUserType(g *gin.Context) {
	var request models.HotelGroupTypeUserType
	g.Bind(&request)

	message, err := repository.DeleteHotelGroupTypeUserType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
