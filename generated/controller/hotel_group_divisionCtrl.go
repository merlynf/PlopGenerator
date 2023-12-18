package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllHotelGroupDivision(g *gin.Context) {
	var request models.HotelGroupDivision
	g.Bind(&request)

	result, err := repository.GetAllHotelGroupDivision(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetHotelGroupDivision(g *gin.Context) {
	var request models.HotelGroupDivision
	g.Bind(&request)

	result, err := repository.GetHotelGroupDivision(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateHotelGroupDivision(g *gin.Context) {
	var request models.HotelGroupDivision
	g.Bind(&request)

	message, err := repository.CreateHotelGroupDivision(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateHotelGroupDivision(g *gin.Context) {
	var request models.HotelGroupDivision
	g.Bind(&request)

	message, err := repository.UpdateHotelGroupDivision(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteHotelGroupDivision(g *gin.Context) {
	var request models.HotelGroupDivision
	g.Bind(&request)

	message, err := repository.DeleteHotelGroupDivision(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
