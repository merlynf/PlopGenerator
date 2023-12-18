package repository

import (
	"golangtemplate/models"
	"log"
	"time"

	uuid "github.com/satori/go.uuid"
)

func GetAllDateOfTheMonth(request models.DateOfTheMonth) ([]models.DateOfTheMonth, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

	var obj []models.DateOfTheMonth

	result := db.Table("date_of_the_month").Scan(&obj)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return obj, err
}

func GetDateOfTheMonth(request models.DateOfTheMonth) (models.DateOfTheMonth, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

	var obj = models.DateOfTheMonth{Id: request.Id}
	result := db.Find(&obj)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return obj, err
}

func CreateDateOfTheMonth(request models.DateOfTheMonth) (string, error) {
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

func UpdateDateOfTheMonth(request models.DateOfTheMonth) (string, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}
    var message string
	var obj = models.DateOfTheMonth{}

	result := db.Model(&obj).Where("id = ?", request.Id).Updates(&request)
    if result.Error == nil {
        message = "data has been successfully updated."
    }

	return message, err
}

func DeleteDateOfTheMonth(request models.DateOfTheMonth) (string, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

    var message string
	var obj = models.DateOfTheMonth{}
	
	result := db.Model(&obj).Where("id = ?", request.Id).Update("status", 0)
    if result.Error == nil {
        message = "data has been successfully deleted."
    }

	return message, err
}


