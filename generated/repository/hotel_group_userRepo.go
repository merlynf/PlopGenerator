package repository

import (
	"golangtemplate/models"
	"log"
	"time"

	uuid "github.com/satori/go.uuid"
)

func GetAllHotelGroupUser(request models.HotelGroupUser) ([]models.HotelGroupUser, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

	var obj []models.HotelGroupUser

	result := db.Table("hotel_group_user").Scan(&obj)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return obj, err
}

func GetHotelGroupUser(request models.HotelGroupUser) (models.HotelGroupUser, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

	var obj = models.HotelGroupUser{Id: request.Id}
	result := db.Find(&obj)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return obj, err
}

func CreateHotelGroupUser(request models.HotelGroupUser) (string, error) {
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

func UpdateHotelGroupUser(request models.HotelGroupUser) (string, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}
    var message string
	var obj = models.HotelGroupUser{}

	result := db.Model(&obj).Where("id = ?", request.Id).Updates(&request)
    if result.Error == nil {
        message = "data has been successfully updated."
    }

	return message, err
}

func DeleteHotelGroupUser(request models.HotelGroupUser) (string, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

    var message string
	var obj = models.HotelGroupUser{}
	
	result := db.Model(&obj).Where("id = ?", request.Id).Update("status", 0)
    if result.Error == nil {
        message = "data has been successfully deleted."
    }

	return message, err
}


