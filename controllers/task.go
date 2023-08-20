package controllers

import (
	"errors"
	"time"
	"myHttpServer/models"
	"myHttpServer/repository"
)

// Save task model to mongoRepo
func SaveTask(username *string, task *models.Task) error {
	err := validateTask(task)
	if err != nil {
		return err
	}
	task.Username = *username
	task.UpdatedAt = time.Now().UnixMilli()
	repository.MongoRepo.SaveTask(task)
	return nil
}


func validateTask(task *models.Task) error {
	if task.Type != models.Music {
		return errors.New("invalid TaskType")
	}
	return nil
}


