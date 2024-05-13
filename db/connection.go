package db

import (
	"database/sql"
	"fmt"
	"task-manager/configs"
)

func OpenConnection() (*sql.DB, error) {
	config := configs.GetDB()

	sc := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", 
		config.Host, config.Port, config.User, config.Pass, config.Database,
	)

	conn, err := sql.Open("postgres", sc)

	if err != nil {
		panic(err)
	}

	err = conn.Ping()

	if err != nil {
		panic(err)
	}

	return conn, err
}