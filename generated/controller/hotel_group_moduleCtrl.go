package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllHotelGroupModule(g *gin.Context) {
	var request models.HotelGroupModule
	g.Bind(&request)

	result, err := repository.GetAllHotelGroupModule(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetHotelGroupModule(g *gin.Context) {
	var request models.HotelGroupModule
	g.Bind(&request)

	result, err := repository.GetHotelGroupModule(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateHotelGroupModule(g *gin.Context) {
	var request models.HotelGroupModule
	g.Bind(&request)

	message, err := repository.CreateHotelGroupModule(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateHotelGroupModule(g *gin.Context) {
	var request models.HotelGroupModule
	g.Bind(&request)

	message, err := repository.UpdateHotelGroupModule(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteHotelGroupModule(g *gin.Context) {
	var request models.HotelGroupModule
	g.Bind(&request)

	message, err := repository.DeleteHotelGroupModule(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
