package dbaccess

import (
	// "database/sql"
	"log"
	_ "github.com/lib/pq"
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	POSTGRES_DB = "postgres"
    POSTGRES_USER =  "postgres"
    POSTGRES_PASSWORD = "postgres"
	POSTGRES_PORT = "5432"
	POSTGRES_HOST = "postgres"
)

const (
	PAGE_SIZE = 20
)

func DB_DSN() string {
	return "postgres://"+POSTGRES_USER+":"+POSTGRES_PASSWORD+"@"+POSTGRES_HOST+":"+POSTGRES_PORT+"/"+ POSTGRES_DB+"?sslmode=disable"
}

func connect() (*sqlx.DB, error) {
	// @TODO: use connection pool
	db, err := sqlx.Connect("postgres", DB_DSN())
	fmt.Println("connect to: ", DB_DSN())
	if err != nil {
		log.Fatal("Failed to open a DB connection: ", err)
	}

	return db, err
}