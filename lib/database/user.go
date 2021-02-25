package database

import (
	"project/config"
	"project/middlewares"
	"project/models"
)

// GetUsers return interface of users
func GetUsers() (interface{}, error) {
	var users []models.Users

	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

// GetUser return interface of a user by given User ID
func GetUser(id int) (interface{}, error) {
	var users []models.Users

	if err := config.DB.Where("id = ?", id).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

// Login return logged in user's data if exist and correct
func Login(user *models.Users) (interface{}, error) {

	if err := config.DB.Where(
		"email = ? AND password = ?", 
		user.Email, 
		user.Password,
	).First(&user).Error; err != nil {
		return nil, err
	}

	// set new user's token
	var err error
	user.Token, err = middlewares.CreateToken(user.ID)
	if err != nil {
		return nil, err
	}

	if err := config.DB.Save(user).Error; err != nil {
		return nil, err
	}

	return  user, nil
}

// Register return successfuly registered data
func Register(user *models.Users) (interface{}, error) {

	if err := config.DB.Save(user).Error; err != nil {
		return nil, err
	}

	// set new user's token
	var err error
	user.Token, err = middlewares.CreateToken(user.ID)
	if err != nil {
		return nil, err
	}

	if err := config.DB.Save(user).Error; err != nil {
		return nil, err
	}

	return user, nil
} 
