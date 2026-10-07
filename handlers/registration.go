package handlers

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RegistrationHandler struct {
	DB *mongo.Database
}

type RegistrationInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (h *RegistrationHandler) RegisterForEvent(c *gin.Context) {
	eventID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid event ID"})
		return
	}

	// check event exists
	var event bson.M
	err = h.DB.Collection("events").FindOne(context.TODO(), bson.M{"_id": eventID}).Decode(&event)
	if err != nil {
		c.JSON(404, gin.H{"error": "Event not found"})
		return
	}

	// check capacity
	registrationCount, _ := h.DB.Collection("registrations").CountDocuments(context.TODO(), bson.M{"event_id": eventID})
	var capacity int64
switch v := event["capacity"].(type) {
case int32:
	capacity = int64(v)
case int64:
	capacity = v
case float64:
	capacity = int64(v)
default:
	c.JSON(500, gin.H{"error": "Event capacity is malformed"})
	return
}
	if registrationCount >= capacity {
		c.JSON(400, gin.H{"error": "Event is full"})
		return
	}

	var input RegistrationInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	// check duplicate
	var existing bson.M
	err = h.DB.Collection("registrations").FindOne(context.TODO(), bson.M{
		"event_id": eventID,
		"email":    input.Email,
	}).Decode(&existing)
	if err == nil {
		c.JSON(400, gin.H{"error": "You are already registered for this event"})
		return
	}

	registration := bson.M{
		"_id":           primitive.NewObjectID(),
		"event_id":      eventID,
		"name":          input.Name,
		"email":         input.Email,
		"registered_at": time.Now(),
	}

	_, err = h.DB.Collection("registrations").InsertOne(context.TODO(), registration)
	if err != nil {
		c.JSON(500, gin.H{"error": "Could not register for event"})
		return
	}

	c.JSON(201, gin.H{"message": "Registration successful", "registration": registration})
}

func (h *RegistrationHandler) GetRegistrations(c *gin.Context) {
	eventID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid event ID"})
		return
	}

	cursor, err := h.DB.Collection("registrations").Find(context.TODO(), bson.M{"event_id": eventID})
	if err != nil {
		c.JSON(500, gin.H{"error": "Could not fetch registrations"})
		return
	}
	defer cursor.Close(context.TODO())

	registrations := []bson.M{}
	if err := cursor.All(context.TODO(), &registrations); err != nil {
		c.JSON(500, gin.H{"error": "Could not parse registrations"})
		return
	}

	c.JSON(200, gin.H{"registrations": registrations})
}