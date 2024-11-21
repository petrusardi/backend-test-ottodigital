package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "postgres"
	dbName = "postgres"
)

func Connect() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	fmt.Println("Database connected!")
}

// migrate -database postgres://postgress:password@localhost:5432/postgres?query -path db/migrations up