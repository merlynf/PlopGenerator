package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllDocumentType(g *gin.Context) {
	var request models.DocumentType
	g.Bind(&request)

	result, err := repository.GetAllDocumentType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetDocumentType(g *gin.Context) {
	var request models.DocumentType
	g.Bind(&request)

	result, err := repository.GetDocumentType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateDocumentType(g *gin.Context) {
	var request models.DocumentType
	g.Bind(&request)

	message, err := repository.CreateDocumentType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateDocumentType(g *gin.Context) {
	var request models.DocumentType
	g.Bind(&request)

	message, err := repository.UpdateDocumentType(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteDocumentType(g *gin.Context) {
	var request models.DocumentType
	g.Bind(&request)

	message, err := repository.DeleteDocumentType(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
