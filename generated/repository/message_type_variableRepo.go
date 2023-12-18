package repository

import (
	"golangtemplate/models"
	"log"
	"time"

	uuid "github.com/satori/go.uuid"
)

func GetAllMessageTypeVariable(request models.MessageTypeVariable) ([]models.MessageTypeVariable, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

	var obj []models.MessageTypeVariable

	result := db.Table("message_type_variable").Scan(&obj)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return obj, err
}

func GetMessageTypeVariable(request models.MessageTypeVariable) (models.MessageTypeVariable, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

	var obj = models.MessageTypeVariable{Id: request.Id}
	result := db.Find(&obj)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return obj, err
}

func CreateMessageTypeVariable(request models.MessageTypeVariable) (string, error) {
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

func UpdateMessageTypeVariable(request models.MessageTypeVariable) (string, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}
    var message string
	var obj = models.MessageTypeVariable{}

	result := db.Model(&obj).Where("id = ?", request.Id).Updates(&request)
    if result.Error == nil {
        message = "data has been successfully updated."
    }

	return message, err
}

func DeleteMessageTypeVariable(request models.MessageTypeVariable) (string, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

    var message string
	var obj = models.MessageTypeVariable{}
	
	result := db.Model(&obj).Where("id = ?", request.Id).Update("status", 0)
    if result.Error == nil {
        message = "data has been successfully deleted."
    }

	return message, err
}


