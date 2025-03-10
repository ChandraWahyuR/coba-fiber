package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB(c Config) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.DB_User,
		c.DB_Password,
		c.DB_Host,
		c.DB_Port,
		c.DB_Name,
	)
	fmt.Println("DSN:", dsn)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println("Error connecting to database:", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Println("Failed to ping database:", err)
		return nil, err
	}

	log.Println("Database connection established successfully.")
	return db, nil
}
