package controllers

import (
	"errors"
	"myHttpServer/models"
	"myHttpServer/repository"
	"time"
)

// Save task model to mongoRepo
func SaveTask(username *string, task *models.Task) error {
	err := validateTask(task)
	transformDate(task)
	if err != nil {
		return err
	}
	task.Username = *username
	task.UpdatedAt = time.Now().UnixMilli()
	repository.MongoRepo.SaveTask(task)
	return nil
}

func SyncTaskToSheet(taskId *string) error {
	task := repository.MongoRepo.GetTaskById(taskId)
	if task == nil {
		return nil
	}
	if task.RowNumber > 0 {
		repository.SheetsRepo.SaveTask(task)
	}
	return nil
}

func validateTask(task *models.Task) error {
	if task.Type != models.Music {
		return errors.New("invalid TaskType")
	}
	return nil
}

func transformDate(task *models.Task) {
	task.Date = time.Unix(task.Date/1000, 0).Truncate(time.Hour * 24).UnixMilli()
}
