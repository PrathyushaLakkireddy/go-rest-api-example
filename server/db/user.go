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

func GetUser(query bson.M, selectObj bson.M) (user models.User, err error) {
	cfg, _ := config.ReadFromFile()
	var c = MongoSession.DB(cfg.MongoDB.Database).C("user")

	err = c.Find(query).Select(selectObj).One(&user)
	return user, err
}

func GetAllUsers(query bson.M, selectObj bson.M) (users []models.User, err error) {
	cfg, _ := config.ReadFromFile()
	var c = MongoSession.DB(cfg.MongoDB.Database).C("user")

	err = c.Find(query).Select(selectObj).All(&users)
	return users, err
}

func DeleteUser(query bson.M) (err error) {
	cfg, _ := config.ReadFromFile()
	var c = MongoSession.DB(cfg.MongoDB.Database).C("user")

	err = c.Remove(query)
	return err
}
