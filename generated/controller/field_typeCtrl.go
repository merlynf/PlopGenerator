package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllFieldType(g *gin.Context) {
	var request models.FieldType
	g.Bind(&request)

	result, err := repository.GetAllFieldType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetFieldType(g *gin.Context) {
	var request models.FieldType
	g.Bind(&request)

	result, err := repository.GetFieldType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateFieldType(g *gin.Context) {
	var request models.FieldType
	g.Bind(&request)

	message, err := repository.CreateFieldType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateFieldType(g *gin.Context) {
	var request models.FieldType
	g.Bind(&request)

	message, err := repository.UpdateFieldType(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteFieldType(g *gin.Context) {
	var request models.FieldType
	g.Bind(&request)

	message, err := repository.DeleteFieldType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
