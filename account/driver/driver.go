package driver

import (
	"fmt"
	"net/url"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/microservice/account/models"
	"github.com/sirupsen/logrus"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func init() {
	if err := godotenv.Load(); err != nil {
		logrus.Error(err)
	}
}

// Config function connect to database
func Config() *gorm.DB {
	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbUser := os.Getenv("DATABASE_USER")
	dbPass := os.Getenv("DATABASE_PASS")
	dbName := os.Getenv("DATABASE_NAME")

	connection := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("sslmode", "disable")
	connStr := fmt.Sprintf("%s?%s", connection, val.Encode())

	connDB, err := gorm.Open("postgres", connStr)
	if err != nil {
		logrus.Error(err)
		return nil
	}

	err = connDB.DB().Ping()
	if err != nil {
		logrus.Error(err)
		return nil
	}

	connDB.Debug().AutoMigrate(
		&models.Account{},
	)

	return connDB
}
