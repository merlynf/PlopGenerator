package repository

import (
	"golangtemplate/models"
	"log"
	"time"

	uuid "github.com/satori/go.uuid"
)

func GetAllEmployeeHistory(request models.EmployeeHistory) ([]models.EmployeeHistory, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

	var obj []models.EmployeeHistory

	result := db.Table("employee_history").Scan(&obj)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return obj, err
}

func GetEmployeeHistory(request models.EmployeeHistory) (models.EmployeeHistory, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

	var obj = models.EmployeeHistory{Id: request.Id}
	result := db.Find(&obj)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return obj, err
}

func CreateEmployeeHistory(request models.EmployeeHistory) (string, error) {
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

func UpdateEmployeeHistory(request models.EmployeeHistory) (string, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}
    var message string
	var obj = models.EmployeeHistory{}

	result := db.Model(&obj).Where("id = ?", request.Id).Updates(&request)
    if result.Error == nil {
        message = "data has been successfully updated."
    }

	return message, err
}

func DeleteEmployeeHistory(request models.EmployeeHistory) (string, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

    var message string
	var obj = models.EmployeeHistory{}
	
	result := db.Model(&obj).Where("id = ?", request.Id).Update("status", 0)
    if result.Error == nil {
        message = "data has been successfully deleted."
    }

	return message, err
}


