package database

import (
	"errors"
	"project/config"
	"project/models"
)

// GetCategories returns all categories stored in database
func GetCategories() (interface{}, error) {
	var cats []models.Categories

	if e := config.DB.Find(&cats).Error; e != nil {
		return nil, e
	}
	return cats, nil
}

// GetCategory returns category by given Category ID
func GetCategory(id string) (interface{}, error) {
	var cats []models.Categories

	if e := config.DB.Where("id = ?", id).Find(&cats).Error; e != nil {
		return nil, e
	}
	return cats, nil
}

// CreateCategories stores new category data to database
func CreateCategories(cat *models.Categories) (interface{}, error) {
	var cats []models.Categories

	if cat.Kode == "" || cat.Nama == "" {
		return nil, errors.New("Kode atau Nama Kategori tidak boleh kosong")
	}

	if e := config.DB.Where("kode = ?", cat.Kode).First(&cats).Error; e != nil {
		if err := config.DB.Save(cat).Error; err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("Kode Kategori sudah ada")
	}

	return cat, nil
}
