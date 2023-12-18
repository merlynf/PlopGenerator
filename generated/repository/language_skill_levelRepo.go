package repository

import (
	"golangtemplate/models"
	"log"
	"time"

	uuid "github.com/satori/go.uuid"
)

func GetAllLanguageSkillLevel(request models.LanguageSkillLevel) ([]models.LanguageSkillLevel, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

	var obj []models.LanguageSkillLevel

	result := db.Table("language_skill_level").Scan(&obj)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return obj, err
}

func GetLanguageSkillLevel(request models.LanguageSkillLevel) (models.LanguageSkillLevel, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

	var obj = models.LanguageSkillLevel{Id: request.Id}
	result := db.Find(&obj)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return obj, err
}

func CreateLanguageSkillLevel(request models.LanguageSkillLevel) (string, error) {
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

func UpdateLanguageSkillLevel(request models.LanguageSkillLevel) (string, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}
    var message string
	var obj = models.LanguageSkillLevel{}

	result := db.Model(&obj).Where("id = ?", request.Id).Updates(&request)
    if result.Error == nil {
        message = "data has been successfully updated."
    }

	return message, err
}

func DeleteLanguageSkillLevel(request models.LanguageSkillLevel) (string, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

    var message string
	var obj = models.LanguageSkillLevel{}
	
	result := db.Model(&obj).Where("id = ?", request.Id).Update("status", 0)
    if result.Error == nil {
        message = "data has been successfully deleted."
    }

	return message, err
}


