package repository

import (
	"golangtemplate/models"
	"log"
	"time"

	uuid "github.com/satori/go.uuid"
)

func GetAllUserTypeModule(request models.UserTypeModule) ([]models.UserTypeModule, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

	var obj []models.UserTypeModule

	result := db.Table("user_type_module").Scan(&obj)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return obj, err
}

func GetUserTypeModule(request models.UserTypeModule) (models.UserTypeModule, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

	var obj = models.UserTypeModule{Id: request.Id}
	result := db.Find(&obj)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return obj, err
}

func CreateUserTypeModule(request models.UserTypeModule) (string, error) {
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

func UpdateUserTypeModule(request models.UserTypeModule) (string, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}
    var message string
	var obj = models.UserTypeModule{}

	result := db.Model(&obj).Where("id = ?", request.Id).Updates(&request)
    if result.Error == nil {
        message = "data has been successfully updated."
    }

	return message, err
}

func DeleteUserTypeModule(request models.UserTypeModule) (string, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

    var message string
	var obj = models.UserTypeModule{}
	
	result := db.Model(&obj).Where("id = ?", request.Id).Update("status", 0)
    if result.Error == nil {
        message = "data has been successfully deleted."
    }

	return message, err
}


