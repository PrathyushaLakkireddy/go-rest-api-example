package main

import (
	"log"

	"github.com/PrathyushaLakkireddy/go-rest-api-example/db"
)

func init() {
	// Initialize db before main to start
	db.InitDB()
}

func main() {
	log.Println("starting main server...")

}
