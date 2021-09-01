package dbaccess

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"
	"fmt"
)

const (
	POSTGRES_DB = "postgres"
    POSTGRES_USER =  "postgres"
    POSTGRES_PASSWORD = "postgres"
	POSTGRES_PORT = "5432"
	POSTGRES_HOST = "localhost"
)

func DB_DSN() string {
	return "postgres://"+POSTGRES_USER+":"+POSTGRES_PASSWORD+"@"+POSTGRES_HOST+":"+POSTGRES_PORT+"/"+ POSTGRES_DB+"?sslmode=disable"
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("postgres", DB_DSN())
	fmt.Println("connect to: ", DB_DSN())
	if err != nil {
		log.Fatal("Failed to open a DB connection: ", err)
	}

	return db, err
}