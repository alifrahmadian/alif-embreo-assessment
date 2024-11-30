package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DB *sql.DB

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func Connect(dbConfig DBConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.DBName)

	DB, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return DB, nil
}
