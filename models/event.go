package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Event struct {
    ID          primitive.ObjectID `bson:"_id,omitempty"`
    Name        string             `bson:"name"`
    Description string             `bson:"description"`
    Date        time.Time          `bson:"date"`
    Location    string             `bson:"location"`
    Capacity    int                `bson:"capacity"`
    Organizer   string             `bson:"organizer"`
    CreatedAt   time.Time          `bson:"created_at"`
}