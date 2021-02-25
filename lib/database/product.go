package database

import (
	"errors"
	"project/config"
	"project/models"
)

// GetProducts return all products and kategori param if any
func GetProducts(kategori string) (interface{}, error) {
	var prods []models.Products

	if e :=  config.DB.Where("categories_id = ?", kategori).Find(&prods).Error; e != nil {
		return nil, e
	}
	return prods, nil
}

// GetProduct return product by given Product ID
func GetProduct(id string) (interface{}, error) {
	var prods []models.Products

	if e := config.DB.Where("id = ?", id).Find(&prods).Error; e != nil {
		return nil, e
	}
	return prods, nil
}

// CreateProducts stores new product data to database
func CreateProducts(prod *models.Products) (interface{}, error) {

	// check if required attributes isn't null or return error
	if requiredNotNull := prod.RequiredNotNull(); requiredNotNull != true {
		return nil, errors.New("Mohon isi data Nama, Stok, Harga, dan Categories ID dengan sesuai")
	} 

	if err := config.DB.Save(prod).Error; err != nil {
		return nil, err
	}
	return prod, nil
}