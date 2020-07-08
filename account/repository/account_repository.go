package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/microservice/account/models"
	"github.com/sirupsen/logrus"
)

//Repository account
type Repository interface {
	CreateNewAccount(data *models.Account) (*models.Account, error)
}

//AccountRepository struct account
type AccountRepository struct {
	DB *gorm.DB
}

//NewAccountRepository implement to service
func NewAccountRepository(db *gorm.DB) Repository {
	return &AccountRepository{
		DB: db,
	}
}

//CreateNewAccount have query create data account to database
func (c *AccountRepository) CreateNewAccount(data *models.Account) (*models.Account, error) {
	if err := c.DB.Table("accounts").Create(&data).Error; err != nil {
		logrus.Error(err)
		return nil, err
	}
	return data, nil
}
