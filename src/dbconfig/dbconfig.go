package dbconfig

import (
	"errors"
	"fmt"
	"os"

	"database/sql"

	_ "github.com/lib/pq"
)

var dbconf DBConfig

type DBConfig struct {
	user     string
	password string
	name     string
	host     string
	port     string
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func setupDB() *sql.DB {
	dbinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbconf.host, dbconf.port, dbconf.user, dbconf.password, dbconf.name)
	db, err := sql.Open("postgres", dbinfo)

	checkErr(err)

	return db
}

func SetDBConfig() error {
	dbconf.user = os.Getenv("DB_USER")
	dbconf.password = os.Getenv("DB_PASSWORD")
	dbconf.name = os.Getenv("DB_NAME")
	dbconf.host = os.Getenv("DB_HOST")
	dbconf.port = os.Getenv("DB_PORT")
	if dbconf.user == "" || dbconf.password == "" || dbconf.name == "" || dbconf.host == "" || dbconf.port == "" {
		return errors.New("One of environment variables is not set: DB_USER, DB_PASSWORD, DB_NAME, DB_HOST, DB_PORT")
	}
	return nil
}
