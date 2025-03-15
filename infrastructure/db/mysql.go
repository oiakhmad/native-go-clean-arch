package db

import (
	"database/sql"
	"fmt"
	"log"

	"native-go-clean-arch/config"

	_ "github.com/go-sql-driver/mysql"
)

// InitMySQL initializes the MySQL database connection
func InitMySQL() *sql.DB {
	connStr := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.Get("DB_USER", "root"),
		config.Get("DB_PASSWORD", "password"),
		config.Get("DB_HOST", "localhost"),
		config.Get("DB_PORT", "3306"),
		config.Get("DB_NAME", "mydatabase"),
	)

	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Database is not reachable:", err)
	}

	log.Println("Connected to MySQL successfully")
	return db
}
