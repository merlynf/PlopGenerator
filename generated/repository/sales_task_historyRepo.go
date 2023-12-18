package repository

import (
	"golangtemplate/models"
	"log"
	"time"

	uuid "github.com/satori/go.uuid"
)

func GetAllSalesTaskHistory(request models.SalesTaskHistory) ([]models.SalesTaskHistory, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

	var obj []models.SalesTaskHistory

	result := db.Table("sales_task_history").Scan(&obj)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return obj, err
}

func GetSalesTaskHistory(request models.SalesTaskHistory) (models.SalesTaskHistory, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

	var obj = models.SalesTaskHistory{Id: request.Id}
	result := db.Find(&obj)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return obj, err
}

func CreateSalesTaskHistory(request models.SalesTaskHistory) (string, error) {
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

func UpdateSalesTaskHistory(request models.SalesTaskHistory) (string, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}
    var message string
	var obj = models.SalesTaskHistory{}

	result := db.Model(&obj).Where("id = ?", request.Id).Updates(&request)
    if result.Error == nil {
        message = "data has been successfully updated."
    }

	return message, err
}

func DeleteSalesTaskHistory(request models.SalesTaskHistory) (string, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

    var message string
	var obj = models.SalesTaskHistory{}
	
	result := db.Model(&obj).Where("id = ?", request.Id).Update("status", 0)
    if result.Error == nil {
        message = "data has been successfully deleted."
    }

	return message, err
}


