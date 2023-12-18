package repository

import (
	"golangtemplate/models"
	"log"
	"time"

	uuid "github.com/satori/go.uuid"
)

func GetAllCustomerAsset(request models.CustomerAsset) ([]models.CustomerAsset, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

	var obj []models.CustomerAsset

	result := db.Table("customer_asset").Scan(&obj)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return obj, err
}

func GetCustomerAsset(request models.CustomerAsset) (models.CustomerAsset, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

	var obj = models.CustomerAsset{Id: request.Id}
	result := db.Find(&obj)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return obj, err
}

func CreateCustomerAsset(request models.CustomerAsset) (string, error) {
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

func UpdateCustomerAsset(request models.CustomerAsset) (string, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}
    var message string
	var obj = models.CustomerAsset{}

	result := db.Model(&obj).Where("id = ?", request.Id).Updates(&request)
    if result.Error == nil {
        message = "data has been successfully updated."
    }

	return message, err
}

func DeleteCustomerAsset(request models.CustomerAsset) (string, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

    var message string
	var obj = models.CustomerAsset{}
	
	result := db.Model(&obj).Where("id = ?", request.Id).Update("status", 0)
    if result.Error == nil {
        message = "data has been successfully deleted."
    }

	return message, err
}


