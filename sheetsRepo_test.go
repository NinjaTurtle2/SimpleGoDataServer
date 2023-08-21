package main

import (
	"fmt"
	"myHttpServer/models"
	"myHttpServer/repository"
	"testing"
)

func TestSheetsService(t *testing.T) {
	fmt.Println("Hello World")
	repository.SheetsRepo.SaveUser(&models.User{
		Username: "yolo",
		RowNumber: 5,
	})
}