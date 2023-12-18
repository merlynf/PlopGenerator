package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllSystemHotelType(g *gin.Context) {
	var request models.SystemHotelType
	g.Bind(&request)

	result, err := repository.GetAllSystemHotelType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetSystemHotelType(g *gin.Context) {
	var request models.SystemHotelType
	g.Bind(&request)

	result, err := repository.GetSystemHotelType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateSystemHotelType(g *gin.Context) {
	var request models.SystemHotelType
	g.Bind(&request)

	message, err := repository.CreateSystemHotelType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateSystemHotelType(g *gin.Context) {
	var request models.SystemHotelType
	g.Bind(&request)

	message, err := repository.UpdateSystemHotelType(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteSystemHotelType(g *gin.Context) {
	var request models.SystemHotelType
	g.Bind(&request)

	message, err := repository.DeleteSystemHotelType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
