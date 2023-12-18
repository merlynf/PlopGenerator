package repository

import (
	"golangtemplate/models"
	"log"
	"time"

	uuid "github.com/satori/go.uuid"
)

func GetAllMessageRecipient(request models.MessageRecipient) ([]models.MessageRecipient, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

	var obj []models.MessageRecipient

	result := db.Table("message_recipient").Scan(&obj)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return obj, err
}

func GetMessageRecipient(request models.MessageRecipient) (models.MessageRecipient, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

	var obj = models.MessageRecipient{Id: request.Id}
	result := db.Find(&obj)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return obj, err
}

func CreateMessageRecipient(request models.MessageRecipient) (string, error) {
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

func UpdateMessageRecipient(request models.MessageRecipient) (string, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}
    var message string
	var obj = models.MessageRecipient{}

	result := db.Model(&obj).Where("id = ?", request.Id).Updates(&request)
    if result.Error == nil {
        message = "data has been successfully updated."
    }

	return message, err
}

func DeleteMessageRecipient(request models.MessageRecipient) (string, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

    var message string
	var obj = models.MessageRecipient{}
	
	result := db.Model(&obj).Where("id = ?", request.Id).Update("status", 0)
    if result.Error == nil {
        message = "data has been successfully deleted."
    }

	return message, err
}


