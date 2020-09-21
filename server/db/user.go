package db

import (
	"log"

	"gopkg.in/mgo.v2/bson"

	"github.com/PrathyushaLakkireddy/go-rest-api-example/config"
	"github.com/PrathyushaLakkireddy/go-rest-api-example/server/modules/user/models"
)

// var cfg *config.Config

func CreateUser(user models.User) error {
	cfg, _ := config.ReadFromFile()
	var c = MongoSession.DB(cfg.MongoDB.Database).C("user")

	err := c.Insert(&user)
	if err != nil {
		log.Println("Error while insertin user details ", err)
	}

	return err
}

func UpdateUser(query bson.M, updateObj bson.M) error {
	cfg, _ := config.ReadFromFile()
	log.Println("cffg..", cfg.MongoDB.Database)
	var c = MongoSession.DB(cfg.MongoDB.Database).C("user")
	err := c.Update(query, updateObj)
	return err
}
