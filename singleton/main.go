package main

import (
	"fmt"
	"singleton/db"
)

func main() {
	dbConn := db.NewDBConnection()
	fmt.Println(dbConn)
	newDBConn := db.NewDBConnection()
	fmt.Println(dbConn)

	fmt.Println(dbConn == newDBConn)
}
