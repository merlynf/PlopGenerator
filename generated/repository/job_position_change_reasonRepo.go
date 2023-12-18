package repository

import (
	"golangtemplate/models"
	"log"
	"time"

	uuid "github.com/satori/go.uuid"
)

func GetAllJobPositionChangeReason(request models.JobPositionChangeReason) ([]models.JobPositionChangeReason, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

	var obj []models.JobPositionChangeReason

	result := db.Table("job_position_change_reason").Scan(&obj)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return obj, err
}

func GetJobPositionChangeReason(request models.JobPositionChangeReason) (models.JobPositionChangeReason, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

	var obj = models.JobPositionChangeReason{Id: request.Id}
	result := db.Find(&obj)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return obj, err
}

func CreateJobPositionChangeReason(request models.JobPositionChangeReason) (string, error) {
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

func UpdateJobPositionChangeReason(request models.JobPositionChangeReason) (string, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}
    var message string
	var obj = models.JobPositionChangeReason{}

	result := db.Model(&obj).Where("id = ?", request.Id).Updates(&request)
    if result.Error == nil {
        message = "data has been successfully updated."
    }

	return message, err
}

func DeleteJobPositionChangeReason(request models.JobPositionChangeReason) (string, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

    var message string
	var obj = models.JobPositionChangeReason{}
	
	result := db.Model(&obj).Where("id = ?", request.Id).Update("status", 0)
    if result.Error == nil {
        message = "data has been successfully deleted."
    }

	return message, err
}


