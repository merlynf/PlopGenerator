package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllSystemHotelGroup(g *gin.Context) {
	var request models.SystemHotelGroup
	g.Bind(&request)

	result, err := repository.GetAllSystemHotelGroup(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetSystemHotelGroup(g *gin.Context) {
	var request models.SystemHotelGroup
	g.Bind(&request)

	result, err := repository.GetSystemHotelGroup(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateSystemHotelGroup(g *gin.Context) {
	var request models.SystemHotelGroup
	g.Bind(&request)

	message, err := repository.CreateSystemHotelGroup(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateSystemHotelGroup(g *gin.Context) {
	var request models.SystemHotelGroup
	g.Bind(&request)

	message, err := repository.UpdateSystemHotelGroup(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteSystemHotelGroup(g *gin.Context) {
	var request models.SystemHotelGroup
	g.Bind(&request)

	message, err := repository.DeleteSystemHotelGroup(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
