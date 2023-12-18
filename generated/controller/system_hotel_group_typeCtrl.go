package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllSystemHotelGroupType(g *gin.Context) {
	var request models.SystemHotelGroupType
	g.Bind(&request)

	result, err := repository.GetAllSystemHotelGroupType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetSystemHotelGroupType(g *gin.Context) {
	var request models.SystemHotelGroupType
	g.Bind(&request)

	result, err := repository.GetSystemHotelGroupType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateSystemHotelGroupType(g *gin.Context) {
	var request models.SystemHotelGroupType
	g.Bind(&request)

	message, err := repository.CreateSystemHotelGroupType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateSystemHotelGroupType(g *gin.Context) {
	var request models.SystemHotelGroupType
	g.Bind(&request)

	message, err := repository.UpdateSystemHotelGroupType(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteSystemHotelGroupType(g *gin.Context) {
	var request models.SystemHotelGroupType
	g.Bind(&request)

	message, err := repository.DeleteSystemHotelGroupType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
