package repository

import (
	"golangtemplate/models"
	"log"
	"time"

	uuid "github.com/satori/go.uuid"
)

func GetAllSystemSalesEventType(request models.SystemSalesEventType) ([]models.SystemSalesEventType, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

	var obj []models.SystemSalesEventType

	result := db.Table("system_sales_event_type").Scan(&obj)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return obj, err
}

func GetSystemSalesEventType(request models.SystemSalesEventType) (models.SystemSalesEventType, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

	var obj = models.SystemSalesEventType{Id: request.Id}
	result := db.Find(&obj)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return obj, err
}

func CreateSystemSalesEventType(request models.SystemSalesEventType) (string, error) {
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

func UpdateSystemSalesEventType(request models.SystemSalesEventType) (string, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}
    var message string
	var obj = models.SystemSalesEventType{}

	result := db.Model(&obj).Where("id = ?", request.Id).Updates(&request)
    if result.Error == nil {
        message = "data has been successfully updated."
    }

	return message, err
}

func DeleteSystemSalesEventType(request models.SystemSalesEventType) (string, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

    var message string
	var obj = models.SystemSalesEventType{}
	
	result := db.Model(&obj).Where("id = ?", request.Id).Update("status", 0)
    if result.Error == nil {
        message = "data has been successfully deleted."
    }

	return message, err
}


