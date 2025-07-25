package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID              primitive.ObjectID `bson:"_id"`
	First_name      *string            `json:"first_name" validate:"required,min=2,max=100"`
	Last_name       *string            `json:"last_name" validate:"required,min=2,max=100"`
	Password        *string            `json:"password" validate:"required,min=6"`
	Email           *string            `json:"email" validate:"required,email"`
	Avatar          *string            `json:"avatar"`
	Phone           *string            `json:"phone" validate:"required"`
	Token           *string            `json:"token"`
	Reference_Token *string            `json:"reference_token"`
	Created_at      time.Time          `json:"created_at"`
	Update_at       time.Time          `json:"update_at"`
	User_id         string             `json:"user_id"`
}
