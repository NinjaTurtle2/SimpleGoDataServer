package models

import (
	"reflect"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Schema struct {
	ID     primitive.ObjectID      `bson:"_id,omitempty"`
	Name   string                  `json:"name"`
	Schema map[string]reflect.Kind `json:"schema"`
}

type User struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Username       string             `json:"username"`
	Password       string             `json:"password"`
	Salt           string             `json:"salt"`
	CreatedAt      int64              `json:"createdAt"`
	UpdatedAt      int64              `json:"updatedAt"`
	LastSessionKey string             `json:"lastSessionKey"`
	RowNumber      int64              `json:"rowNumber"`
}

type Session struct {
	Key        string `json:"key"`
	Value      string `json:"value"`
	CreatedAt  int64  `json:"createdAt"`
	ValidUntil int64  `json:"validUntil"`
}

type Data struct {
	ID        primitive.ObjectID     `bson:"_id,omitempty"`
	SchemaID  string                 `json:"schemaId"`
	Data      map[string]interface{} `json:"data"`
	CreatedAt int64                  `json:"createdAt"`
}

type Task struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Username  string             `json:"username"`
	Date      int64              `json:"date"`
	Type      TaskType           `json:"type"`
	Complete  bool               `json:"complete"`
	Duration  int64              `json:"duration"`
	UpdatedAt int64              `json:"updatedAt"`
	RowNumber int64              `json:"rowNumber"`
}

type TaskType string

const (
	Music TaskType = "MUSIC"
)
