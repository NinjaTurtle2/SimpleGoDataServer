package controllers

import (
	"errors"
	"reflect"
	"time"
	"myHttpServer/models"
	"myHttpServer/repository"
)

// Save data model to mongoRepo
func SaveData(data *models.Data) error {
	err := validateData(data)
	if err != nil {
		return err
	}
	data.CreatedAt = time.Now().UnixMilli()
	repository.MongoRepo.SaveData(data)
	return nil
}


func validateData(data *models.Data) error {
	if data.SchemaID == "" {
		return errors.New("schemaId is required")
	}
	schema := getCachedSchema(data.SchemaID)
	if schema == nil {
		return errors.New("schema not found")
	}
	for key, value := range data.Data {
		typeOfValue := reflect.TypeOf(value)
		if schema.Schema[key] != typeOfValue.Kind() {
			return errors.New("data type mismatch " + key + " expected " + schema.Schema[key].String() + " got " + typeOfValue.Kind().String())
		}
	}
	return nil
}


func getCachedSchema(id string) *models.Schema {
	//Fetch from redis if not found fetch from mongo and cache it in redis
	schema, _ := repository.RedisRepo.GetSchemaByID(id)
	if schema == nil {
		schema = repository.MongoRepo.GetSchemaByID(id)
		if schema != nil {
			repository.RedisRepo.SaveSchema(schema)
		}
	}
	return schema
}

