package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllHotelGroupTypeJobTitle(g *gin.Context) {
	var request models.HotelGroupTypeJobTitle
	g.Bind(&request)

	result, err := repository.GetAllHotelGroupTypeJobTitle(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetHotelGroupTypeJobTitle(g *gin.Context) {
	var request models.HotelGroupTypeJobTitle
	g.Bind(&request)

	result, err := repository.GetHotelGroupTypeJobTitle(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateHotelGroupTypeJobTitle(g *gin.Context) {
	var request models.HotelGroupTypeJobTitle
	g.Bind(&request)

	message, err := repository.CreateHotelGroupTypeJobTitle(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateHotelGroupTypeJobTitle(g *gin.Context) {
	var request models.HotelGroupTypeJobTitle
	g.Bind(&request)

	message, err := repository.UpdateHotelGroupTypeJobTitle(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteHotelGroupTypeJobTitle(g *gin.Context) {
	var request models.HotelGroupTypeJobTitle
	g.Bind(&request)

	message, err := repository.DeleteHotelGroupTypeJobTitle(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
