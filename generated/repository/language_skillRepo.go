package repository

import (
	"golangtemplate/models"
	"log"
	"time"

	uuid "github.com/satori/go.uuid"
)

func GetAllLanguageSkill(request models.LanguageSkill) ([]models.LanguageSkill, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

	var obj []models.LanguageSkill

	result := db.Table("language_skill").Scan(&obj)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return obj, err
}

func GetLanguageSkill(request models.LanguageSkill) (models.LanguageSkill, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

	var obj = models.LanguageSkill{Id: request.Id}
	result := db.Find(&obj)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return obj, err
}

func CreateLanguageSkill(request models.LanguageSkill) (string, error) {
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

func UpdateLanguageSkill(request models.LanguageSkill) (string, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}
    var message string
	var obj = models.LanguageSkill{}

	result := db.Model(&obj).Where("id = ?", request.Id).Updates(&request)
    if result.Error == nil {
        message = "data has been successfully updated."
    }

	return message, err
}

func DeleteLanguageSkill(request models.LanguageSkill) (string, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

    var message string
	var obj = models.LanguageSkill{}
	
	result := db.Model(&obj).Where("id = ?", request.Id).Update("status", 0)
    if result.Error == nil {
        message = "data has been successfully deleted."
    }

	return message, err
}


