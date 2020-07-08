package service

import (
	"github.com/microservice/product/models"
	"github.com/microservice/product/repository"
	"github.com/sirupsen/logrus"
)

// Service interface
type Service interface {
	CreateNewProduct(data *models.Product) (*models.Product, error)
}

// ProductService struct
type ProductService struct {
	ProductRepository repository.Repository
}

// NewProductService implement to handler
func NewProductService(productRepository repository.Repository) Service {
	return &ProductService{
		ProductRepository: productRepository,
	}
}

//CreateNewProduct use to create data product
func (c *ProductService) CreateNewProduct(data *models.Product) (*models.Product, error) {
	productResponse, errProduct := c.ProductRepository.CreateNewProduct(data)
	if errProduct != nil {
		logrus.Error(errProduct)
	}
	return productResponse, nil
}
