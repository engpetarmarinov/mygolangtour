package database

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
	"time"
)

var DbConn *sql.DB

// Connection parameters
var (
	server = "localhost"
	port   = 1433
	//TODO: take credentials from env variables
	user     = "mygolangtour"
	password = "mygolangtour"
	database = "mygolangtour.rest.products"
)

// SetupDatabase
func SetupDatabase() {
	// Create connection string
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s",
		server, user, password, port, database)

	var err error
	DbConn, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal(err)
	}

	DbConn.SetMaxOpenConns(3)
	DbConn.SetMaxIdleConns(3)
	DbConn.SetConnMaxLifetime(60 * time.Second)
}
