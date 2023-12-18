package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllHotelGroupUserType(g *gin.Context) {
	var request models.HotelGroupUserType
	g.Bind(&request)

	result, err := repository.GetAllHotelGroupUserType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetHotelGroupUserType(g *gin.Context) {
	var request models.HotelGroupUserType
	g.Bind(&request)

	result, err := repository.GetHotelGroupUserType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateHotelGroupUserType(g *gin.Context) {
	var request models.HotelGroupUserType
	g.Bind(&request)

	message, err := repository.CreateHotelGroupUserType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateHotelGroupUserType(g *gin.Context) {
	var request models.HotelGroupUserType
	g.Bind(&request)

	message, err := repository.UpdateHotelGroupUserType(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteHotelGroupUserType(g *gin.Context) {
	var request models.HotelGroupUserType
	g.Bind(&request)

	message, err := repository.DeleteHotelGroupUserType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
