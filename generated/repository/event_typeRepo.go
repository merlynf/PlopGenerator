package repository

import (
	"golangtemplate/models"
	"log"
	"time"

	uuid "github.com/satori/go.uuid"
)

func GetAllEventType(request models.EventType) ([]models.EventType, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

	var obj []models.EventType

	result := db.Table("event_type").Scan(&obj)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return obj, err
}

func GetEventType(request models.EventType) (models.EventType, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

	var obj = models.EventType{Id: request.Id}
	result := db.Find(&obj)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return obj, err
}

func CreateEventType(request models.EventType) (string, error) {
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

func UpdateEventType(request models.EventType) (string, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}
    var message string
	var obj = models.EventType{}

	result := db.Model(&obj).Where("id = ?", request.Id).Updates(&request)
    if result.Error == nil {
        message = "data has been successfully updated."
    }

	return message, err
}

func DeleteEventType(request models.EventType) (string, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

    var message string
	var obj = models.EventType{}
	
	result := db.Model(&obj).Where("id = ?", request.Id).Update("status", 0)
    if result.Error == nil {
        message = "data has been successfully deleted."
    }

	return message, err
}

