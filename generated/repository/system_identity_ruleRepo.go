package repository

import (
	"golangtemplate/models"
	"log"
	"time"

	uuid "github.com/satori/go.uuid"
)

func GetAllSystemIdentityRule(request models.SystemIdentityRule) ([]models.SystemIdentityRule, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

	var obj []models.SystemIdentityRule

	result := db.Table("system_identity_rule").Scan(&obj)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return obj, err
}

func GetSystemIdentityRule(request models.SystemIdentityRule) (models.SystemIdentityRule, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

	var obj = models.SystemIdentityRule{Id: request.Id}
	result := db.Find(&obj)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return obj, err
}

func CreateSystemIdentityRule(request models.SystemIdentityRule) (string, error) {
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

func UpdateSystemIdentityRule(request models.SystemIdentityRule) (string, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}
    var message string
	var obj = models.SystemIdentityRule{}

	result := db.Model(&obj).Where("id = ?", request.Id).Updates(&request)
    if result.Error == nil {
        message = "data has been successfully updated."
    }

	return message, err
}

func DeleteSystemIdentityRule(request models.SystemIdentityRule) (string, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

    var message string
	var obj = models.SystemIdentityRule{}
	
	result := db.Model(&obj).Where("id = ?", request.Id).Update("status", 0)
    if result.Error == nil {
        message = "data has been successfully deleted."
    }

	return message, err
}


