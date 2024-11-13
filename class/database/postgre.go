package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func PgConnect(dbName string, dbUser string, dbPassword string) (*sql.DB, error) {
	return sql.Open("postgres", fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=localhost", dbUser, dbPassword, dbName))
}

func DbOpen(dbName string) *sql.DB {
	db, err := PgConnect(dbName, "postgres", "postgres")
	if err != nil {
		log.Fatal(err)
	}

	return db
}
