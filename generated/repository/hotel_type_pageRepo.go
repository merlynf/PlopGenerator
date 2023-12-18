package repository

import (
	"golangtemplate/models"
	"log"
	"time"

	uuid "github.com/satori/go.uuid"
)

func GetAllHotelTypePage(request models.HotelTypePage) ([]models.HotelTypePage, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

	var obj []models.HotelTypePage

	result := db.Table("hotel_type_page").Scan(&obj)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return obj, err
}

func GetHotelTypePage(request models.HotelTypePage) (models.HotelTypePage, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

	var obj = models.HotelTypePage{Id: request.Id}
	result := db.Find(&obj)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return obj, err
}

func CreateHotelTypePage(request models.HotelTypePage) (string, error) {
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

func UpdateHotelTypePage(request models.HotelTypePage) (string, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}
    var message string
	var obj = models.HotelTypePage{}

	result := db.Model(&obj).Where("id = ?", request.Id).Updates(&request)
    if result.Error == nil {
        message = "data has been successfully updated."
    }

	return message, err
}

func DeleteHotelTypePage(request models.HotelTypePage) (string, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

    var message string
	var obj = models.HotelTypePage{}
	
	result := db.Model(&obj).Where("id = ?", request.Id).Update("status", 0)
    if result.Error == nil {
        message = "data has been successfully deleted."
    }

	return message, err
}


