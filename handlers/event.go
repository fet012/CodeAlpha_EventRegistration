package handlers

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type EventHandler struct {
	DB *mongo.Database
}

type EventInput struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	Location    string    `json:"location"`
	Capacity    int       `json:"capacity"`
	Organizer   string    `json:"organizer"`
}

func (h *EventHandler) CreateEvent(c *gin.Context) {
	var input EventInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	event := bson.M{
		"_id":         primitive.NewObjectID(),
		"name":        input.Name,
		"description": input.Description,
		"date":        input.Date,
		"location":    input.Location,
		"capacity":    input.Capacity,
		"organizer":   input.Organizer,
		"created_at":  time.Now(),
	}

	_, err := h.DB.Collection("events").InsertOne(context.TODO(), event)
	if err != nil {
		c.JSON(500, gin.H{"error": "Could not create event"})
		return
	}

	c.JSON(201, gin.H{"message": "Event created successfully", "event": event})
}

func (h *EventHandler) ListEvents(c *gin.Context) {
	cursor, err := h.DB.Collection("events").Find(context.TODO(), bson.M{})
	if err != nil {
		c.JSON(500, gin.H{"error": "Could not fetch events"})
		return
	}
	defer cursor.Close(context.TODO())

	var events []bson.M
	if err := cursor.All(context.TODO(), &events); err != nil {
		c.JSON(500, gin.H{"error": "Could not parse events"})
		return
	}

	c.JSON(200, gin.H{"events": events})
}

func (h *EventHandler) GetEvent(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid event ID"})
		return
	}

	var event bson.M
	err = h.DB.Collection("events").FindOne(context.TODO(), bson.M{"_id": id}).Decode(&event)
	if err != nil {
		c.JSON(404, gin.H{"error": "Event not found"})
		return
	}

	c.JSON(200, gin.H{"event": event})
}