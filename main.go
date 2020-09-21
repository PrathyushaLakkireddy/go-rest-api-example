package main

import (
	"log"
	"net/http"

	"github.com/rs/cors"

	"github.com/PrathyushaLakkireddy/go-rest-api-example/config"
	"github.com/PrathyushaLakkireddy/go-rest-api-example/server"
	"github.com/PrathyushaLakkireddy/go-rest-api-example/server/db"
)

func main() {
	cfg, err := config.ReadFromFile()
	if err != nil {
		log.Fatal(err)
	}

	db.InitDB(cfg)
	defer db.MongoSession.Close()

	log.Println("starting main server...")

	router := server.Router()

	methods := []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodOptions, http.MethodPut}
	origins := []string{"*"}

	c := cors.New(cors.Options{
		AllowedOrigins:   origins,
		AllowedMethods:   methods,
		AllowCredentials: true,
		MaxAge:           1000,
		AllowedHeaders:   []string{"*"},
	})

	log.Fatal(http.ListenAndServe(":"+cfg.ServerPort, c.Handler(router)))
}
