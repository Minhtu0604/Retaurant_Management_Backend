package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Food struct {
	ID         primitive.ObjectID `bson:"_id" `
	Name       *string            `json:"name" validate:"required,min=2,max=100"`
	Price      *float64           `json:"price" validate:"required"`
	Food_image *string            `json:"food_image" validate:"required"`
	Created_at time.Time          `json:"created_at"`
	Update_at  time.Time          `json:"update_at"`
	Food_id    *string            `json:"food_id"`
	Menu_id    *string            `json:"menu_id" validate:"required"`
}
