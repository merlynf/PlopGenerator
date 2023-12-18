package repository

import (
	"golangtemplate/models"
	"log"
	"time"

	uuid "github.com/satori/go.uuid"
)

func GetAllRoundNumberFunction(request models.RoundNumberFunction) ([]models.RoundNumberFunction, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

	var obj []models.RoundNumberFunction

	result := db.Table("round_number_function").Scan(&obj)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return obj, err
}

func GetRoundNumberFunction(request models.RoundNumberFunction) (models.RoundNumberFunction, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

	var obj = models.RoundNumberFunction{Id: request.Id}
	result := db.Find(&obj)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return obj, err
}

func CreateRoundNumberFunction(request models.RoundNumberFunction) (string, error) {
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

func UpdateRoundNumberFunction(request models.RoundNumberFunction) (string, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}
    var message string
	var obj = models.RoundNumberFunction{}

	result := db.Model(&obj).Where("id = ?", request.Id).Updates(&request)
    if result.Error == nil {
        message = "data has been successfully updated."
    }

	return message, err
}

func DeleteRoundNumberFunction(request models.RoundNumberFunction) (string, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

    var message string
	var obj = models.RoundNumberFunction{}
	
	result := db.Model(&obj).Where("id = ?", request.Id).Update("status", 0)
    if result.Error == nil {
        message = "data has been successfully deleted."
    }

	return message, err
}


