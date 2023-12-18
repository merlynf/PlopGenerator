package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllHotelTeam(g *gin.Context) {
	var request models.HotelTeam
	g.Bind(&request)

	result, err := repository.GetAllHotelTeam(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetHotelTeam(g *gin.Context) {
	var request models.HotelTeam
	g.Bind(&request)

	result, err := repository.GetHotelTeam(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateHotelTeam(g *gin.Context) {
	var request models.HotelTeam
	g.Bind(&request)

	message, err := repository.CreateHotelTeam(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateHotelTeam(g *gin.Context) {
	var request models.HotelTeam
	g.Bind(&request)

	message, err := repository.UpdateHotelTeam(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteHotelTeam(g *gin.Context) {
	var request models.HotelTeam
	g.Bind(&request)

	message, err := repository.DeleteHotelTeam(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
