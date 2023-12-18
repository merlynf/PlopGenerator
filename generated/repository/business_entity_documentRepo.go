package repository

import (
	"golangtemplate/models"
	"log"
	"time"

	uuid "github.com/satori/go.uuid"
)

func GetAllBusinessEntityDocument(request models.BusinessEntityDocument) ([]models.BusinessEntityDocument, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

	var obj []models.BusinessEntityDocument

	result := db.Table("business_entity_document").Scan(&obj)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return obj, err
}

func GetBusinessEntityDocument(request models.BusinessEntityDocument) (models.BusinessEntityDocument, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

	var obj = models.BusinessEntityDocument{Id: request.Id}
	result := db.Find(&obj)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return obj, err
}

func CreateBusinessEntityDocument(request models.BusinessEntityDocument) (string, error) {
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

func UpdateBusinessEntityDocument(request models.BusinessEntityDocument) (string, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}
    var message string
	var obj = models.BusinessEntityDocument{}

	result := db.Model(&obj).Where("id = ?", request.Id).Updates(&request)
    if result.Error == nil {
        message = "data has been successfully updated."
    }

	return message, err
}

func DeleteBusinessEntityDocument(request models.BusinessEntityDocument) (string, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

    var message string
	var obj = models.BusinessEntityDocument{}
	
	result := db.Model(&obj).Where("id = ?", request.Id).Update("status", 0)
    if result.Error == nil {
        message = "data has been successfully deleted."
    }

	return message, err
}


