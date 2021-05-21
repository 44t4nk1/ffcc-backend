package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Course struct {
	ID      primitive.ObjectID `json:"_id" bson:"_id"`
	Code    string             `json:"code" bson:"code"`
	Credits int                `json:"credits" bson:"credits"`
	Faculty string             `json:"faculty" bson:"faculty"`
	Owner   string             `json:"owner" bson:"owner"`
	Room    string             `json:"room" bson:"room"`
	Slot    string             `json:"slot" bson:"slot"`
	Title   string             `json:"title" bson:"title"`
	Type    string             `json:"type" bson:"type"`
}