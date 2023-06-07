package controllers

import (
	"myHttpServer/models"
	"myHttpServer/repository"
)
func GetSchemas() []models.Schema {
	return repository.MongoRepo.GetSchemas()
}

func GetSchemaByID(id string) *models.Schema {
	return repository.MongoRepo.GetSchemaByID(id)
}

func SaveSchema(schema *models.Schema) {
	//Call MongoRepo to save schema
	repository.MongoRepo.SaveSchema(schema)
}