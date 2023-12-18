package repository

import (
	"golangtemplate/models"
	"log"
	"time"

	uuid "github.com/satori/go.uuid"
)

func GetAllTermCategoryTerm(request models.TermCategoryTerm) ([]models.TermCategoryTerm, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

	var obj []models.TermCategoryTerm

	result := db.Table("term_category_term").Scan(&obj)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return obj, err
}

func GetTermCategoryTerm(request models.TermCategoryTerm) (models.TermCategoryTerm, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

	var obj = models.TermCategoryTerm{Id: request.Id}
	result := db.Find(&obj)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return obj, err
}

func CreateTermCategoryTerm(request models.TermCategoryTerm) (string, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

    var message string

	request.Id = (uuid.NewV4()).String()
	request.CreatedAt = time.Now()
    
	result := db.Create(&request)
    if result.Error == nil {
        message = "data has been successfully added."
    }

	return message, err
}

func UpdateTermCategoryTerm(request models.TermCategoryTerm) (string, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}
    var message string
	var obj = models.TermCategoryTerm{}

	result := db.Model(&obj).Where("id = ?", request.Id).Updates(&request)
    if result.Error == nil {
        message = "data has been successfully updated."
    }

	return message, err
}

func DeleteTermCategoryTerm(request models.TermCategoryTerm) (string, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

    var message string
	var obj = models.TermCategoryTerm{}
	
	result := db.Model(&obj).Where("id = ?", request.Id).Update("status", 0)
    if result.Error == nil {
        message = "data has been successfully deleted."
    }

	return message, err
}


