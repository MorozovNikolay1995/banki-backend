package dbconfig

import (
	"errors"
	"os"
)

type DBConfig struct {
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	DB_HOST     string
	DB_PORT     string
}

func GetDBConfig() (*DBConfig, error) {

	var dbconf DBConfig
	dbconf.DB_USER = os.Getenv("DB_USER")
	dbconf.DB_PASSWORD = os.Getenv("DB_PASSWORD")
	dbconf.DB_NAME = os.Getenv("DB_NAME")
	dbconf.DB_HOST = os.Getenv("DB_HOST")
	dbconf.DB_PORT = os.Getenv("DB_PORT")
	if dbconf.DB_USER == "" || dbconf.DB_PASSWORD == "" || dbconf.DB_NAME == "" || dbconf.DB_HOST == "" || dbconf.DB_PORT == "" {
		return nil, errors.New("One of environment variables is not set: DB_USER, DB_PASSWORD, DB_NAME, DB_HOST, DB_PORT")
	}

	return &dbconf, nil
}
