package repository

import (
	"golangtemplate/models"
	"log"
	"time"

	uuid "github.com/satori/go.uuid"
)

func GetAllEmployeeAdditionalRole(request models.EmployeeAdditionalRole) ([]models.EmployeeAdditionalRole, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

	var obj []models.EmployeeAdditionalRole

	result := db.Table("employee_additional_role").Scan(&obj)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return obj, err
}

func GetEmployeeAdditionalRole(request models.EmployeeAdditionalRole) (models.EmployeeAdditionalRole, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

	var obj = models.EmployeeAdditionalRole{Id: request.Id}
	result := db.Find(&obj)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return obj, err
}

func CreateEmployeeAdditionalRole(request models.EmployeeAdditionalRole) (string, error) {
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

func UpdateEmployeeAdditionalRole(request models.EmployeeAdditionalRole) (string, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}
    var message string
	var obj = models.EmployeeAdditionalRole{}

	result := db.Model(&obj).Where("id = ?", request.Id).Updates(&request)
    if result.Error == nil {
        message = "data has been successfully updated."
    }

	return message, err
}

func DeleteEmployeeAdditionalRole(request models.EmployeeAdditionalRole) (string, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

    var message string
	var obj = models.EmployeeAdditionalRole{}
	
	result := db.Model(&obj).Where("id = ?", request.Id).Update("status", 0)
    if result.Error == nil {
        message = "data has been successfully deleted."
    }

	return message, err
}


