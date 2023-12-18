package repository

import (
	"golangtemplate/models"
	"log"
	"time"

	uuid "github.com/satori/go.uuid"
)

func GetAllEmailTemplateForm(request models.EmailTemplateForm) ([]models.EmailTemplateForm, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

	var obj []models.EmailTemplateForm

	result := db.Table("email_template_form").Scan(&obj)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return obj, err
}

func GetEmailTemplateForm(request models.EmailTemplateForm) (models.EmailTemplateForm, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

	var obj = models.EmailTemplateForm{Id: request.Id}
	result := db.Find(&obj)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return obj, err
}

func CreateEmailTemplateForm(request models.EmailTemplateForm) (string, error) {
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

func UpdateEmailTemplateForm(request models.EmailTemplateForm) (string, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}
    var message string
	var obj = models.EmailTemplateForm{}

	result := db.Model(&obj).Where("id = ?", request.Id).Updates(&request)
    if result.Error == nil {
        message = "data has been successfully updated."
    }

	return message, err
}

func DeleteEmailTemplateForm(request models.EmailTemplateForm) (string, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

    var message string
	var obj = models.EmailTemplateForm{}
	
	result := db.Model(&obj).Where("id = ?", request.Id).Update("status", 0)
    if result.Error == nil {
        message = "data has been successfully deleted."
    }

	return message, err
}


