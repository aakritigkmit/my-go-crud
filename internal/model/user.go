package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name    string             `bson:"name,omitempty"`
	Age     int                `bson:"age,omitempty"`
	Country string             `bson:"country,omitempty"`
}
