package controller

import (
	"net/http"
	"golangtemplate/models"
	"golangtemplate/repository"

	"github.com/gin-gonic/gin"
)

func GetAllTermCategoryTerm(g *gin.Context) {
	var request models.TermCategoryTerm
	g.Bind(&request)

	result, err := repository.GetAllTermCategoryTerm(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func GetTermCategoryTerm(g *gin.Context) {
	var request models.TermCategoryTerm
	g.Bind(&request)

	result, err := repository.GetTermCategoryTerm(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "data": result})
	}
}

func CreateTermCategoryTerm(g *gin.Context) {
	var request models.TermCategoryTerm
	g.Bind(&request)

	message, err := repository.CreateTermCategoryTerm(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func UpdateTermCategoryTerm(g *gin.Context) {
	var request models.TermCategoryTerm
	g.Bind(&request)

	message, err := repository.UpdateTermCategoryTerm(request)

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}

func DeleteTermCategoryTerm(g *gin.Context) {
	var request models.TermCategoryTerm
	g.Bind(&request)

	message, err := repository.DeleteTermCategoryTerm(request)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "No Record."})
	} else {
		g.JSON(http.StatusOK, gin.H{"success": true, "message": message})
	}
}
