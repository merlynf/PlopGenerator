package repository

import (
	"golangtemplate/models"
	"log"
	"time"

	uuid "github.com/satori/go.uuid"
)

func GetAllSalesProspectRevenue(request models.SalesProspectRevenue) ([]models.SalesProspectRevenue, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

	var obj []models.SalesProspectRevenue

	result := db.Table("sales_prospect_revenue").Scan(&obj)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return obj, err
}

func GetSalesProspectRevenue(request models.SalesProspectRevenue) (models.SalesProspectRevenue, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

	var obj = models.SalesProspectRevenue{Id: request.Id}
	result := db.Find(&obj)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return obj, err
}

func CreateSalesProspectRevenue(request models.SalesProspectRevenue) (string, error) {
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

func UpdateSalesProspectRevenue(request models.SalesProspectRevenue) (string, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}
    var message string
	var obj = models.SalesProspectRevenue{}

	result := db.Model(&obj).Where("id = ?", request.Id).Updates(&request)
    if result.Error == nil {
        message = "data has been successfully updated."
    }

	return message, err
}

func DeleteSalesProspectRevenue(request models.SalesProspectRevenue) (string, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

    var message string
	var obj = models.SalesProspectRevenue{}
	
	result := db.Model(&obj).Where("id = ?", request.Id).Update("status", 0)
    if result.Error == nil {
        message = "data has been successfully deleted."
    }

	return message, err
}


