package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllHotelCustomer(g *gin.Context) {
	var request models.HotelCustomer
	g.Bind(&request)

	result, err := repository.GetAllHotelCustomer(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetHotelCustomer(g *gin.Context) {
	var request models.HotelCustomer
	g.Bind(&request)

	result, err := repository.GetHotelCustomer(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateHotelCustomer(g *gin.Context) {
	var request models.HotelCustomer
	g.Bind(&request)

	message, err := repository.CreateHotelCustomer(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateHotelCustomer(g *gin.Context) {
	var request models.HotelCustomer
	g.Bind(&request)

	message, err := repository.UpdateHotelCustomer(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteHotelCustomer(g *gin.Context) {
	var request models.HotelCustomer
	g.Bind(&request)

	message, err := repository.DeleteHotelCustomer(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
