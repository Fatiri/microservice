package service

import (
	"github.com/microservice/account/models"
	"github.com/microservice/account/repository"
	"github.com/sirupsen/logrus"
)

//Service account interface
type Service interface {
	CreateNewAccount(data *models.Account) (*models.Account, error)
}

//AccountService struct
type AccountService struct {
	AccountRepository repository.Repository
}

//NewAccountService func account service registration
func NewAccountService(accountRepository repository.Repository) Service {
	return &AccountService{
		AccountRepository: accountRepository,
	}
}

//CreateNewAccount have query create data account to database
func (c *AccountService) CreateNewAccount(data *models.Account) (*models.Account, error) {
	accountResponse, errAccount := c.AccountRepository.CreateNewAccount(data)
	if errAccount != nil {
		logrus.Error(errAccount)
	}
	return accountResponse, nil
}
