package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllHotelGroupCustomer(g *gin.Context) {
	var request models.HotelGroupCustomer
	g.Bind(&request)

	result, err := repository.GetAllHotelGroupCustomer(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetHotelGroupCustomer(g *gin.Context) {
	var request models.HotelGroupCustomer
	g.Bind(&request)

	result, err := repository.GetHotelGroupCustomer(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateHotelGroupCustomer(g *gin.Context) {
	var request models.HotelGroupCustomer
	g.Bind(&request)

	message, err := repository.CreateHotelGroupCustomer(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateHotelGroupCustomer(g *gin.Context) {
	var request models.HotelGroupCustomer
	g.Bind(&request)

	message, err := repository.UpdateHotelGroupCustomer(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteHotelGroupCustomer(g *gin.Context) {
	var request models.HotelGroupCustomer
	g.Bind(&request)

	message, err := repository.DeleteHotelGroupCustomer(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
