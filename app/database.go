package app

import (
	"database/sql"
	"time"

	"github.com/andreaswiidi/restful-api/helper"
	_ "github.com/lib/pq" // add this
)

func NewDB() *sql.DB {
	connStr := "user=postgres password=123456 dbname=databaseTest sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
