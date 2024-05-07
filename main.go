package main

import (
	"log"

	"github.com/krisfragata/flow-api/db"
)

func main(){
	Database, err := db.New()
	if err != nil {
		log.Fatal("Error connecting to DB:", err)
	}
	log.Printf("Last posted data: %+v", Database)
	server := NewServer(":3000")
	server.Run()
}
