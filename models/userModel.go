package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id       primitive.ObjectID `bson:"_id"`
	Name     string             `json:"name" validate:"min=2, max=10"`
	Email    string             `json :"email" validate:"email, required"`
	Password string             `json :"password" validate:"required"`
}
