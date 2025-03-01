package main

import (
	"backend/db"
	"fmt"
	"log"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatal("Could not connect to database: ", err)
	}
	fmt.Println("Database connected: ", database)

}