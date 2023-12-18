package repository

import (
	"golangtemplate/models"
	"log"
	"time"

	uuid "github.com/satori/go.uuid"
)

func GetAllHotelTypeUserType(request models.HotelTypeUserType) ([]models.HotelTypeUserType, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

	var obj []models.HotelTypeUserType

	result := db.Table("hotel_type_user_type").Scan(&obj)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return obj, err
}

func GetHotelTypeUserType(request models.HotelTypeUserType) (models.HotelTypeUserType, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

	var obj = models.HotelTypeUserType{Id: request.Id}
	result := db.Find(&obj)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return obj, err
}

func CreateHotelTypeUserType(request models.HotelTypeUserType) (string, error) {
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

func UpdateHotelTypeUserType(request models.HotelTypeUserType) (string, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}
    var message string
	var obj = models.HotelTypeUserType{}

	result := db.Model(&obj).Where("id = ?", request.Id).Updates(&request)
    if result.Error == nil {
        message = "data has been successfully updated."
    }

	return message, err
}

func DeleteHotelTypeUserType(request models.HotelTypeUserType) (string, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

    var message string
	var obj = models.HotelTypeUserType{}
	
	result := db.Model(&obj).Where("id = ?", request.Id).Update("status", 0)
    if result.Error == nil {
        message = "data has been successfully deleted."
    }

	return message, err
}


