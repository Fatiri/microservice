package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/microservice/product/models"
	"github.com/sirupsen/logrus"
)

// Repository interface
type Repository interface {
	CreateNewProduct(data *models.Product) (*models.Product, error)
}

// ProductRepository struct
type ProductRepository struct {
	DB *gorm.DB
}

// NewProductRepository implement to service
func NewProductRepository(db *gorm.DB) Repository {
	return &ProductRepository{
		DB: db,
	}
}

//CreateNewProduct create new data product
func (c *ProductRepository) CreateNewProduct(data *models.Product) (*models.Product, error) {
	if err := c.DB.Table("products").Create(&data).Error; err != nil {
		logrus.Error(err)
		return nil, err
	}
	return data, nil
}
