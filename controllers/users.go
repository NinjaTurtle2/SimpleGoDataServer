package controllers

import (
	"myHttpServer/models"
	"myHttpServer/repository"
	"myHttpServer/utils"
	"time"
)

//Get User using mongoRepo
func GetUser(id string) (*models.User,error) {
	user := repository.MongoRepo.GetUserByID(id)
	if user == nil {
		return nil, &utils.NotFoundError{Message: "User not found"}
	}
	return user, nil
}

//Get User by username using mongoRepo
func GetUserByUsername(username string) (*models.User, error) {
	user := repository.MongoRepo.GetUserByUsername(username)
	if user == nil {
		return nil, &utils.NotFoundError{Message: "User not found"}
	}
	return user, nil
}

//Get Users using mongoRepo
func GetUsers() []models.User {
	return repository.MongoRepo.GetUsers()
}


//Create User using mongoRepo	
func CreateUser(user *models.User) (*models.User, error) {
	
	existingUser, _ := GetUserByUsername(user.Username)
	if existingUser != nil {
		return nil, &utils.AlreadyExistsError{Message: "User already exists"}
	}

	//Set CreatedAt
	user.CreatedAt = time.Now().UnixMilli()
	user.UpdatedAt = time.Now().UnixMilli()
	user.Salt = utils.GenerateRandomSalt()
	user.Password = utils.HashPassword(user.Password, user.Salt)
	repository.MongoRepo.SaveUser(user)
	return user, nil
}

//Delete User Using mongoRepo
func DeleteUser(id string) error {
	repository.MongoRepo.DeleteUser(id)
	return nil
}

//Login User using mongoRepo
func LoginUser(user *models.User) (*models.Session, error) {
	existingUser, err := GetUserByUsername(user.Username)

	if err != nil {
		return nil, err
	}

	//Hash user password
	user.Password = utils.HashPassword(user.Password, existingUser.Salt)
	if user.Password != existingUser.Password {
		return nil, &utils.InvalidCredentialsError{Message: "Invalid credentials"}
	}
	if existingUser.LastSessionKey != "" {
		DeleteSession(existingUser.LastSessionKey)
	}

	session, err := CreateSession(existingUser)
	if err != nil {
		return nil, err
	}

	existingUser.LastSessionKey = session.Key
	existingUser.UpdatedAt = time.Now().UnixMilli()
	updateUser(existingUser)

	return session, nil
}


func updateUser(user *models.User) (*models.User, error) {
	repository.MongoRepo.UpdateUser(user)
	return user, nil
}
