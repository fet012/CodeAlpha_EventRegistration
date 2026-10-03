package handlers

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LoginInput struct{
	Email string `json:"email"`
	Password string `json:"password"`
}

type EventInput struct {
    ID          primitive.ObjectID `json:"_id,omitempty"`
    Name        string             `json:"name"`
    Description string             `json:"description"`
    Date        time.Time          `json:"date"`
    Location    string             `json:"location"`
    Capacity    int                `json:"capacity"`
    Organizer   string             `json:"organizer"`
    CreatedAt   time.Time          `json:"created_at"`
}
func login(c *gin.Context){
	godotenv.Load()
	var input LoginInput
	c.ShouldBindJSON(&input)

	if input.Email == os.Getenv("ADMIN_EMAIL") && input.Password == os.Getenv("ADMIN_PASSWORD"){
		token, _ := generateToken(input.Email)
		c.JSON(200, gin.H{"token": token})
		return
	}
	c.JSON(401, gin.H{"error": "Invalid credentials"})
}

func CreateEvent(c *gin.Context){
	var input EventInput
	c.ShouldBindJSON((&input))

	c.JSON(201, gin.H{EventInput})


	
}
