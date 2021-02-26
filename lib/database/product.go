package database

import (
	"errors"
	"project/config"
	"project/models"
)

// GetProducts return all products stored in database
func GetProducts() (interface{}, error) {
	var prods []models.Products

	if e := config.DB.Find(&prods).Error; e != nil {
		return nil, e
	}
	return prods, nil
}

// GetProductsWithCategory return all products and kategori param if any
func GetProductsWithCategory(kategori string) (interface{}, error) {
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
	var cats []models.Categories

	// check if required attributes isn't null or return error
	if requiredNotNull := prod.RequiredNotNull(); requiredNotNull != true {
		return nil, errors.New("Mohon isi data Nama, Stok, Harga, dan Categories ID dengan sesuai")
	}

	if e := config.DB.Where("kode = ?", prod.CategoriesID).First(&cats).Error; e != nil {
		return nil, errors.New("Kode Kategori belum ada")
	} else {
		if err := config.DB.Save(prod).Error; err != nil {
			return nil, err
		}
	}

	return prod, nil
}
