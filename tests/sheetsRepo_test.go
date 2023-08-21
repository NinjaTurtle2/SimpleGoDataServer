package tests

import (
	"fmt"
	"myHttpServer/models"
	"myHttpServer/repository"
	"testing"
)

func TestSheetsService(t *testing.T) {
	fmt.Println("Hello World")
	repository.SheetsRepo.SaveTask(&models.Task{
		Username: "yolo",
		Type: models.Music,
	})
}