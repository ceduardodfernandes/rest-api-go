package models

import (
    "gopkg.in/mgo.v2/bson"
)

type User struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	name 		string `bson:"name" json:"name"`
	gravatar 	string `bson:"gravatar" json:"gravatar"`
}

type Widget struct {
	ID 			bson.ObjectId `bson:"_id" json:"id"`
	name		string `bson:"name" json:"name"`
	color		string `bson:"color" json:"color"`
	price		string `bson:"price" json:"price"`
	inventory	int `bson:"inventory" json:"inventory"`
	melts		bool `bson:"melts" json:"melts"`
}
