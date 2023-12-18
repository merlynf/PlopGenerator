package repository

import (
	"golangtemplate/models"
	"log"
	"time"

	uuid "github.com/satori/go.uuid"
)

func GetAllRoundNumberDigit(request models.RoundNumberDigit) ([]models.RoundNumberDigit, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

	var obj []models.RoundNumberDigit

	result := db.Table("round_number_digit").Scan(&obj)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return obj, err
}

func GetRoundNumberDigit(request models.RoundNumberDigit) (models.RoundNumberDigit, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

	var obj = models.RoundNumberDigit{Id: request.Id}
	result := db.Find(&obj)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return obj, err
}

func CreateRoundNumberDigit(request models.RoundNumberDigit) (string, error) {
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

func UpdateRoundNumberDigit(request models.RoundNumberDigit) (string, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}
    var message string
	var obj = models.RoundNumberDigit{}

	result := db.Model(&obj).Where("id = ?", request.Id).Updates(&request)
    if result.Error == nil {
        message = "data has been successfully updated."
    }

	return message, err
}

func DeleteRoundNumberDigit(request models.RoundNumberDigit) (string, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

    var message string
	var obj = models.RoundNumberDigit{}
	
	result := db.Model(&obj).Where("id = ?", request.Id).Update("status", 0)
    if result.Error == nil {
        message = "data has been successfully deleted."
    }

	return message, err
}


