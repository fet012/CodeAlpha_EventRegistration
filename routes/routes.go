package routes

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	
	"EventRegistration/handlers"
	"EventRegistration/middleware"
)

func SetupRoutes(r *gin.Engine, db *mongo.Database) {
	eventHandler := &handlers.EventHandler{DB: db}
	registrationHandler := &handlers.RegistrationHandler{DB: db}

	// public routes
	r.POST("/admin/login", handlers.Login)
	r.GET("/events/:id", eventHandler.GetEvent)
	r.POST("/events/:id/register", registrationHandler.RegisterForEvent)

	// protected routes
	admin := r.Group("/")
	admin.Use(middleware.AuthMiddleware())
	{
		admin.POST("/events", eventHandler.CreateEvent)
		admin.GET("/events", eventHandler.ListEvents)
		admin.GET("/events/:id/registrations", registrationHandler.GetRegistrations)
	}
}