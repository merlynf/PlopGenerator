package repository

import (
	"golangtemplate/models"
	"log"
	"time"

	uuid "github.com/satori/go.uuid"
)

func GetAllMessageVariable(request models.MessageVariable) ([]models.MessageVariable, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

	var obj []models.MessageVariable

	result := db.Table("message_variable").Scan(&obj)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return obj, err
}

func GetMessageVariable(request models.MessageVariable) (models.MessageVariable, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

	var obj = models.MessageVariable{Id: request.Id}
	result := db.Find(&obj)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return obj, err
}

func CreateMessageVariable(request models.MessageVariable) (string, error) {
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

func UpdateMessageVariable(request models.MessageVariable) (string, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}
    var message string
	var obj = models.MessageVariable{}

	result := db.Model(&obj).Where("id = ?", request.Id).Updates(&request)
    if result.Error == nil {
        message = "data has been successfully updated."
    }

	return message, err
}

func DeleteMessageVariable(request models.MessageVariable) (string, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

    var message string
	var obj = models.MessageVariable{}
	
	result := db.Model(&obj).Where("id = ?", request.Id).Update("status", 0)
    if result.Error == nil {
        message = "data has been successfully deleted."
    }

	return message, err
}

