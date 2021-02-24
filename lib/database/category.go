package database

import (
	"project/config"
	"project/models"
)

func GetCategories() (interface{}, error) {
	var cats []models.Categories

	if e := config.DB.Find(&cats).Error; e != nil {
		return nil, e
	}
	return cats, nil
}

func GetCategory(id string) (interface{}, error) {
	var cats []models.Categories

	if e := config.DB.Where("id = ?", id).Find(&cats).Error; e != nil {
		return nil, e
	}
	return cats, nil
}

func CreateCategories(cat *models.Categories) (interface{}, error) {

	if err := config.DB.Save(cat).Error; err != nil {
		return nil, err
	}
	return cat, nil
}
