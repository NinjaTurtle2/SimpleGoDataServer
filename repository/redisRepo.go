package repository

import (
	"encoding/json"
	"fmt"
	"myHttpServer/models"
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
	"myHttpServer/utils"
)

// redis global variable
var RedisRepo = NewRedisRepo()

// RedisRepo is a struct that holds the redis client
type redisRepo struct {
	client          *redis.Client
	sessionDuration int64
}

// NewRedisRepo returns a new instance of RedisRepo
func NewRedisRepo() *redisRepo {
	godotenv.Load()
	//redisUsername := os.Getenv("REDIS_USERNAME")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: redisPassword,
		DB:       0, // use default DB
	})
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	sessionDuration, _ := strconv.ParseInt(os.Getenv("SESSION_DURATION_MINUTES"), 10, 64)
	return &redisRepo{client, sessionDuration * 60}
}

// GetSessionByKey returns a session by key
func (r *redisRepo) GetSessionByKey(key string) (*models.Session, error) {
	val, err := r.client.Get(utils.SESSION_KEY_PREFIX + key).Result()
	if err != nil {
		return nil, err
	}
	var session models.Session
	err = json.Unmarshal([]byte(val), &session)
	if err != nil {
		return nil, err
	}
	return &session, nil
}

// SaveSession saves a session
func (r *redisRepo) SaveSession(session *models.Session) error {
	session.ValidUntil = session.CreatedAt + r.sessionDuration*1000
	value, err := json.Marshal(session)
	if err != nil {
		return err
	}
	status := r.client.Set(utils.SESSION_KEY_PREFIX+session.Key, value, time.Duration(session.ValidUntil))
	if status.Err() != nil {
		return status.Err()
	}
	return nil
}

// DeleteSession deletes a session
func (r *redisRepo) DeleteSession(key string) error {
	status := r.client.Del(utils.SESSION_KEY_PREFIX + key)
	if status.Err() != nil {
		return status.Err()
	}
	return nil
}

// GetSchemaById returns a schema by id
func (r *redisRepo) GetSchemaByID(id string) (*models.Schema, error) {
	val, err := r.client.Get(utils.SCHEMA_KEY_PREFIX + id).Result()
	if err != nil {
		return nil, err
	}
	var schema models.Schema
	err = json.Unmarshal([]byte(val), &schema)
	if err != nil {
		return nil, err
	}
	return &schema, nil
}

// SaveSchema saves a schema
func (r *redisRepo) SaveSchema(schema *models.Schema) error {
	value, err := json.Marshal(schema)
	if err != nil {
		return err
	}
	status := r.client.Set(utils.SCHEMA_KEY_PREFIX+schema.ID.Hex(), value, 0)
	if status.Err() != nil {
		return status.Err()
	}
	return nil
}
