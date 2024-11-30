package db

import (
	"database/sql"
	"fmt"

	config "github.com/alifrahmadian/alif-embreo-assessment/configs"
)

var DB *sql.DB

func Connect(dbConfig *config.DBConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.DBName)

	DB, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return DB, nil
}
