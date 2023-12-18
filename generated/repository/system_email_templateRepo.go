package repository

import (
	"golangtemplate/models"
	"log"
	"time"

	uuid "github.com/satori/go.uuid"
)

func GetAllSystemEmailTemplate(request models.SystemEmailTemplate) ([]models.SystemEmailTemplate, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

	var obj []models.SystemEmailTemplate

	result := db.Table("system_email_template").Scan(&obj)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return obj, err
}

func GetSystemEmailTemplate(request models.SystemEmailTemplate) (models.SystemEmailTemplate, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

	var obj = models.SystemEmailTemplate{Id: request.Id}
	result := db.Find(&obj)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return obj, err
}

func CreateSystemEmailTemplate(request models.SystemEmailTemplate) (string, error) {
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

func UpdateSystemEmailTemplate(request models.SystemEmailTemplate) (string, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}
    var message string
	var obj = models.SystemEmailTemplate{}

	result := db.Model(&obj).Where("id = ?", request.Id).Updates(&request)
    if result.Error == nil {
        message = "data has been successfully updated."
    }

	return message, err
}

func DeleteSystemEmailTemplate(request models.SystemEmailTemplate) (string, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

    var message string
	var obj = models.SystemEmailTemplate{}
	
	result := db.Model(&obj).Where("id = ?", request.Id).Update("status", 0)
    if result.Error == nil {
        message = "data has been successfully deleted."
    }

	return message, err
}


