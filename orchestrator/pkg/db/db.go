package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func NewDB(config *Config) (*sql.DB, error) {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", config.Host, 5432, config.User, config.Password, config.DBName)
	dbCon, err := sql.Open(config.Driver, connectionString)
	if err != nil {
		return nil, err
	}
	return dbCon, nil
}
