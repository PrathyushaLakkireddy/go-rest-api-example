package db

import (
	"log"
	"time"

	"gopkg.in/mgo.v2"

	"github.com/PrathyushaLakkireddy/go-rest-api-example/config"
)

var cfg *config.Config

var (
	MongoDbUrl = &mgo.DialInfo{
		Addrs:    []string{string("localhost")},
		Timeout:  30 * time.Second,
		Username: cfg.MongoDB.Username,
		Password: cfg.MongoDB.Password,
		Database: cfg.MongoDB.Database,
	}
)

func InitDB() {
	MongoSession, err := mgo.DialWithInfo(MongoDbUrl)

	if err != nil {
		log.Fatalf("Error while connecting to database", err)
	}

	if err = MongoSession.Ping(); err != nil {
		defer MongoSession.Close()
		log.Fatalf("Error while connecting to Database ", err)
	}

	log.Println("Database connected successfully")
}
