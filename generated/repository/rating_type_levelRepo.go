package repository

import (
	"golangtemplate/models"
	"log"
	"time"

	uuid "github.com/satori/go.uuid"
)

func GetAllRatingTypeLevel(request models.RatingTypeLevel) ([]models.RatingTypeLevel, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

	var obj []models.RatingTypeLevel

	result := db.Table("rating_type_level").Scan(&obj)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return obj, err
}

func GetRatingTypeLevel(request models.RatingTypeLevel) (models.RatingTypeLevel, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

	var obj = models.RatingTypeLevel{Id: request.Id}
	result := db.Find(&obj)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return obj, err
}

func CreateRatingTypeLevel(request models.RatingTypeLevel) (string, error) {
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

func UpdateRatingTypeLevel(request models.RatingTypeLevel) (string, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}
    var message string
	var obj = models.RatingTypeLevel{}

	result := db.Model(&obj).Where("id = ?", request.Id).Updates(&request)
    if result.Error == nil {
        message = "data has been successfully updated."
    }

	return message, err
}

func DeleteRatingTypeLevel(request models.RatingTypeLevel) (string, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

    var message string
	var obj = models.RatingTypeLevel{}
	
	result := db.Model(&obj).Where("id = ?", request.Id).Update("status", 0)
    if result.Error == nil {
        message = "data has been successfully deleted."
    }

	return message, err
}


