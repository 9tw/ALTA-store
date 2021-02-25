package database

import (
	"project/config"
	"project/models"
)

func GetProducts(kategori string) (interface{}, error) {
	var prods []models.Products

	if e :=  config.DB.Where("id_kategori = ?", kategori).Find(&prods).Error; e != nil {
		return nil, e
	}
	return prods, nil
}

func GetProduct(id string) (interface{}, error) {
	var prods []models.Products

	if e := config.DB.Where("id = ?", id).Find(&prods).Error; e != nil {
		return nil, e
	}
	return prods, nil
}

func CreateProducts(prod *models.Products) (interface{}, error) {

	if err := config.DB.Save(prod).Error; err != nil {
		return nil, err
	}
	return prod, nil
}