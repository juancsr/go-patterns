package db

import "fmt"

type dbConnection struct {
	Username string
	Password string
	URL      string
}

var PSQLConnection *dbConnection

func NewDBConnection() *dbConnection {
	if PSQLConnection == nil {
		fmt.Println("Initiate a new PSQLConnection")
		PSQLConnection = &dbConnection{Username: "test", Password: "test", URL: "my.test.url.com"}
	}
	return PSQLConnection
}
