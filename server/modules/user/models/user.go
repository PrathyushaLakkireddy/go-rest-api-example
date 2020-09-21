package models

import "gopkg.in/mgo.v2/bson"

type (
	User struct {
		ID       bson.ObjectId `json:"_id" bson:"_id,omitempty"`
		Name     string        `json:"name" bson:"name"`
		Password string        `json:"password" bson:"password"`
		Email    string        `json:"email" bson:"email"`
	}
)
