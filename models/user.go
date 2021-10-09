package models

import "gopkg.in/mgo.v2/bson"

type User struct{
	Id         bson.ObjectId    `json:"id"  bson:"_id"`
	Name	   string			`json:"name"  bson:"name"`
	Email 	   string			`json:"email"  bson:"email"`
	Password   string			`json:"password"  bson:"password"`
}

type Post struct{
	Id         bson.ObjectId    `json:"id"  bson:"_id"`
	userId 	   string			`json:"userId"  bson:"userId"`
	Caption	   string			`json:"caption"  bson:"caption"`
	Image 	   string			`json:"image"  bson:"image"`
	Timestamp  string			`json:"timestamp"  bson:"_createdAt"`
}

// https://docs.google.com/document/d/1sFhVumoczf_PmaL_R__Rm9AHqaHsUWgj1x9YcQP6Is4/preview?pru=AAABfIQlzWc*SNh3nqvVDkBiMz-6n7hm4g#