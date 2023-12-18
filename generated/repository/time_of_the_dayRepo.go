package repository

import (
	"golangtemplate/models"
	"log"
	"time"

	uuid "github.com/satori/go.uuid"
)

func GetAllTimeOfTheDay(request models.TimeOfTheDay) ([]models.TimeOfTheDay, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

	var obj []models.TimeOfTheDay

	result := db.Table("time_of_the_day").Scan(&obj)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return obj, err
}

func GetTimeOfTheDay(request models.TimeOfTheDay) (models.TimeOfTheDay, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

	var obj = models.TimeOfTheDay{Id: request.Id}
	result := db.Find(&obj)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return obj, err
}

func CreateTimeOfTheDay(request models.TimeOfTheDay) (string, error) {
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

func UpdateTimeOfTheDay(request models.TimeOfTheDay) (string, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}
    var message string
	var obj = models.TimeOfTheDay{}

	result := db.Model(&obj).Where("id = ?", request.Id).Updates(&request)
    if result.Error == nil {
        message = "data has been successfully updated."
    }

	return message, err
}

func DeleteTimeOfTheDay(request models.TimeOfTheDay) (string, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

    var message string
	var obj = models.TimeOfTheDay{}
	
	result := db.Model(&obj).Where("id = ?", request.Id).Update("status", 0)
    if result.Error == nil {
        message = "data has been successfully deleted."
    }

	return message, err
}


