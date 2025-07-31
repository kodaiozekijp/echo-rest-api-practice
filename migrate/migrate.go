package main

import (
	"echo-rest-api-practice/db"
	"echo-rest-api-practice/entities"
	"fmt"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&entities.User{}, &entities.Task{})
}
