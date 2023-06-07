package controllers

import (
	"myHttpServer/models"
	"myHttpServer/repository"
	"myHttpServer/utils"
	"time"
)

// GetSession by key from redisrepo
func GetSession(key string) (*models.Session, error) {
	session, err := repository.RedisRepo.GetSessionByKey(key)
	if err != nil {
		return nil, err
	}
	if session == nil {
		return nil, &utils.NotFoundError{Message: "Session not found"}
	}
	
	return session, nil
}

//Save session to redisrepo
func CreateSession(user *models.User) (*models.Session, error) {
	var session models.Session
	session.Key = utils.GenerateRandomSalt()
	session.Value = user.ID.Hex()
	session.CreatedAt = time.Now().UnixMilli()
	err:= repository.RedisRepo.SaveSession(&session)
	if err != nil {
		return nil, err
	}
	return &session, nil
}

//Delete session from redisrepo
func DeleteSession(key string) error {
	err := repository.RedisRepo.DeleteSession(key)
	if err != nil {
		return err
	}
	return nil
}



