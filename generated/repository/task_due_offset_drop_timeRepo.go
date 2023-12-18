package repository

import (
	"golangtemplate/models"
	"log"
	"time"

	uuid "github.com/satori/go.uuid"
)

func GetAllTaskDueOffsetDropTime(request models.TaskDueOffsetDropTime) ([]models.TaskDueOffsetDropTime, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

	var obj []models.TaskDueOffsetDropTime

	result := db.Table("task_due_offset_drop_time").Scan(&obj)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	return obj, err
}

func GetTaskDueOffsetDropTime(request models.TaskDueOffsetDropTime) (models.TaskDueOffsetDropTime, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

	var obj = models.TaskDueOffsetDropTime{Id: request.Id}
	result := db.Find(&obj)
	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return obj, err
}

func CreateTaskDueOffsetDropTime(request models.TaskDueOffsetDropTime) (string, error) {
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

func UpdateTaskDueOffsetDropTime(request models.TaskDueOffsetDropTime) (string, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}
    var message string
	var obj = models.TaskDueOffsetDropTime{}

	result := db.Model(&obj).Where("id = ?", request.Id).Updates(&request)
    if result.Error == nil {
        message = "data has been successfully updated."
    }

	return message, err
}

func DeleteTaskDueOffsetDropTime(request models.TaskDueOffsetDropTime) (string, error) {
	db, err := models.Connectdb()
	if err != nil {
		panic(err)
	}

    var message string
	var obj = models.TaskDueOffsetDropTime{}
	
	result := db.Model(&obj).Where("id = ?", request.Id).Update("status", 0)
    if result.Error == nil {
        message = "data has been successfully deleted."
    }

	return message, err
}


