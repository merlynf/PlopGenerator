package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllSystemHotel(g *gin.Context) {
	var request models.SystemHotel
	g.Bind(&request)

	result, err := repository.GetAllSystemHotel(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetSystemHotel(g *gin.Context) {
	var request models.SystemHotel
	g.Bind(&request)

	result, err := repository.GetSystemHotel(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateSystemHotel(g *gin.Context) {
	var request models.SystemHotel
	g.Bind(&request)

	message, err := repository.CreateSystemHotel(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateSystemHotel(g *gin.Context) {
	var request models.SystemHotel
	g.Bind(&request)

	message, err := repository.UpdateSystemHotel(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteSystemHotel(g *gin.Context) {
	var request models.SystemHotel
	g.Bind(&request)

	message, err := repository.DeleteSystemHotel(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
