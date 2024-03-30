package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Design struct {
	Id        primitive.ObjectID `json:"id" bson:"_id"`
	Visible   bool               `json:"-" bson:"visible"`
	Path      string             `json:"path" bson:"path"`
	Extension string             `json:"extension" bson:"extension"`
}
