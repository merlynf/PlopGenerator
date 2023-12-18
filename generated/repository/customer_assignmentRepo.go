package repository

import (
	"golangtemplate/models"
	"log"
	"time"

	uuid "github.com/satori/go.uuid"
)

func GetAllCustomerAssignment(request models.CustomerAssignment) ([]models.CustomerAssignment, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

	var obj []models.CustomerAssignment

	result := db.Table("customer_assignment").Scan(&obj)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return obj, err
}

func GetCustomerAssignment(request models.CustomerAssignment) (models.CustomerAssignment, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

	var obj = models.CustomerAssignment{Id: request.Id}
	result := db.Find(&obj)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return obj, err
}

func CreateCustomerAssignment(request models.CustomerAssignment) (string, error) {
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

func UpdateCustomerAssignment(request models.CustomerAssignment) (string, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}
    var message string
	var obj = models.CustomerAssignment{}

	result := db.Model(&obj).Where("id = ?", request.Id).Updates(&request)
    if result.Error == nil {
        message = "data has been successfully updated."
    }

	return message, err
}

func DeleteCustomerAssignment(request models.CustomerAssignment) (string, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

    var message string
	var obj = models.CustomerAssignment{}
	
	result := db.Model(&obj).Where("id = ?", request.Id).Update("status", 0)
    if result.Error == nil {
        message = "data has been successfully deleted."
    }

	return message, err
}


