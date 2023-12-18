package repository

import (
	"golangtemplate/models"
	"log"
	"time"

	uuid "github.com/satori/go.uuid"
)

func GetAllAdditionalDetailType(request models.AdditionalDetailType) ([]models.AdditionalDetailType, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

	var obj []models.AdditionalDetailType

	result := db.Table("additional_detail_type").Scan(&obj)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return obj, err
}

func GetAdditionalDetailType(request models.AdditionalDetailType) (models.AdditionalDetailType, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

	var obj = models.AdditionalDetailType{Id: request.Id}
	result := db.Find(&obj)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return obj, err
}

func CreateAdditionalDetailType(request models.AdditionalDetailType) (string, error) {
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

func UpdateAdditionalDetailType(request models.AdditionalDetailType) (string, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}
    var message string
	var obj = models.AdditionalDetailType{}

	result := db.Model(&obj).Where("id = ?", request.Id).Updates(&request)
    if result.Error == nil {
        message = "data has been successfully updated."
    }

	return message, err
}

func DeleteAdditionalDetailType(request models.AdditionalDetailType) (string, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

    var message string
	var obj = models.AdditionalDetailType{}
	
	result := db.Model(&obj).Where("id = ?", request.Id).Update("status", 0)
    if result.Error == nil {
        message = "data has been successfully deleted."
    }

	return message, err
}


