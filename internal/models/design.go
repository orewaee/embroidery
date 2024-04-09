package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Design struct {
	Id        primitive.ObjectID `json:"id" bson:"_id"`
	Name      string             `json:"name" bson:"name"`
	Tag       string             `json:"tag" bson:"tag"`
	Visible   bool               `json:"-" bson:"visible"`
	Path      string             `json:"-" bson:"path"`
	Extension string             `json:"-" bson:"extension"`
}
