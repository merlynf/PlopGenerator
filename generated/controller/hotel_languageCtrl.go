package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllHotelLanguage(g *gin.Context) {
	var request models.HotelLanguage
	g.Bind(&request)

	result, err := repository.GetAllHotelLanguage(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetHotelLanguage(g *gin.Context) {
	var request models.HotelLanguage
	g.Bind(&request)

	result, err := repository.GetHotelLanguage(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateHotelLanguage(g *gin.Context) {
	var request models.HotelLanguage
	g.Bind(&request)

	message, err := repository.CreateHotelLanguage(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateHotelLanguage(g *gin.Context) {
	var request models.HotelLanguage
	g.Bind(&request)

	message, err := repository.UpdateHotelLanguage(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteHotelLanguage(g *gin.Context) {
	var request models.HotelLanguage
	g.Bind(&request)

	message, err := repository.DeleteHotelLanguage(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
