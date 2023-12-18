package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllHotelTypeUserType(g *gin.Context) {
	var request models.HotelTypeUserType
	g.Bind(&request)

	result, err := repository.GetAllHotelTypeUserType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetHotelTypeUserType(g *gin.Context) {
	var request models.HotelTypeUserType
	g.Bind(&request)

	result, err := repository.GetHotelTypeUserType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateHotelTypeUserType(g *gin.Context) {
	var request models.HotelTypeUserType
	g.Bind(&request)

	message, err := repository.CreateHotelTypeUserType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateHotelTypeUserType(g *gin.Context) {
	var request models.HotelTypeUserType
	g.Bind(&request)

	message, err := repository.UpdateHotelTypeUserType(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteHotelTypeUserType(g *gin.Context) {
	var request models.HotelTypeUserType
	g.Bind(&request)

	message, err := repository.DeleteHotelTypeUserType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
