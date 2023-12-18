package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllHotelGroupPage(g *gin.Context) {
	var request models.HotelGroupPage
	g.Bind(&request)

	result, err := repository.GetAllHotelGroupPage(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetHotelGroupPage(g *gin.Context) {
	var request models.HotelGroupPage
	g.Bind(&request)

	result, err := repository.GetHotelGroupPage(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateHotelGroupPage(g *gin.Context) {
	var request models.HotelGroupPage
	g.Bind(&request)

	message, err := repository.CreateHotelGroupPage(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateHotelGroupPage(g *gin.Context) {
	var request models.HotelGroupPage
	g.Bind(&request)

	message, err := repository.UpdateHotelGroupPage(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteHotelGroupPage(g *gin.Context) {
	var request models.HotelGroupPage
	g.Bind(&request)

	message, err := repository.DeleteHotelGroupPage(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
