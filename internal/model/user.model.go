package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name      string             `bson:"name,omitempty" validate:"required,min=2,max=100"`
	Email     string             `bson:"email,omitempty" validate:"required,email"`
	Password  string             `bson:"password,omitempty" `
	Age       int                `bson:"age,omitempty" validate:"required,gte=18,lte=100"`
	Country   string             `bson:"country,omitempty"`
	CreatedAt time.Time          `bson:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
}
