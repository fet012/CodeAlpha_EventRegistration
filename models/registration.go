package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Register struct {
	ID   primitive.ObjectID  `bson:"_id,omitempty"`
	Name string `bson:"name"`
	Email string `bson:"email"`
	Event primitive.ObjectID `bson:"eventId"`
	RegisterAt time.Time `bson:"registeredAt"`
}