package repository

import (
	"golangtemplate/models"
	"log"
	"time"

	uuid "github.com/satori/go.uuid"
)

func GetAllHotelHotelRevenueSource(request models.HotelHotelRevenueSource) ([]models.HotelHotelRevenueSource, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

	var obj []models.HotelHotelRevenueSource

	result := db.Table("hotel_hotel_revenue_source").Scan(&obj)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return obj, err
}

func GetHotelHotelRevenueSource(request models.HotelHotelRevenueSource) (models.HotelHotelRevenueSource, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

	var obj = models.HotelHotelRevenueSource{Id: request.Id}
	result := db.Find(&obj)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return obj, err
}

func CreateHotelHotelRevenueSource(request models.HotelHotelRevenueSource) (string, error) {
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

func UpdateHotelHotelRevenueSource(request models.HotelHotelRevenueSource) (string, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}
    var message string
	var obj = models.HotelHotelRevenueSource{}

	result := db.Model(&obj).Where("id = ?", request.Id).Updates(&request)
    if result.Error == nil {
        message = "data has been successfully updated."
    }

	return message, err
}

func DeleteHotelHotelRevenueSource(request models.HotelHotelRevenueSource) (string, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

    var message string
	var obj = models.HotelHotelRevenueSource{}
	
	result := db.Model(&obj).Where("id = ?", request.Id).Update("status", 0)
    if result.Error == nil {
        message = "data has been successfully deleted."
    }

	return message, err
}


