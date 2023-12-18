package repository

import (
	"golangtemplate/models"
	"log"
	"time"

	uuid "github.com/satori/go.uuid"
)

func GetAllUserSocialMediaAccount(request models.UserSocialMediaAccount) ([]models.UserSocialMediaAccount, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

	var obj []models.UserSocialMediaAccount

	result := db.Table("user_social_media_account").Scan(&obj)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return obj, err
}

func GetUserSocialMediaAccount(request models.UserSocialMediaAccount) (models.UserSocialMediaAccount, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

	var obj = models.UserSocialMediaAccount{Id: request.Id}
	result := db.Find(&obj)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return obj, err
}

func CreateUserSocialMediaAccount(request models.UserSocialMediaAccount) (string, error) {
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

func UpdateUserSocialMediaAccount(request models.UserSocialMediaAccount) (string, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}
    var message string
	var obj = models.UserSocialMediaAccount{}

	result := db.Model(&obj).Where("id = ?", request.Id).Updates(&request)
    if result.Error == nil {
        message = "data has been successfully updated."
    }

	return message, err
}

func DeleteUserSocialMediaAccount(request models.UserSocialMediaAccount) (string, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

    var message string
	var obj = models.UserSocialMediaAccount{}
	
	result := db.Model(&obj).Where("id = ?", request.Id).Update("status", 0)
    if result.Error == nil {
        message = "data has been successfully deleted."
    }

	return message, err
}


