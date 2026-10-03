package handlers

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func generateToken(email string) (string, error) {
	godotenv.Load()
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func Login(c *gin.Context) {
	godotenv.Load()
	var input LoginInput
	c.ShouldBindJSON(&input)

	if input.Email == os.Getenv("ADMIN_EMAIL") && input.Password == os.Getenv("ADMIN_PASSWORD") {
		token, err := generateToken(input.Email)
		if err != nil {
			c.JSON(500, gin.H{"error": "Could not generate token"})
			return
		}
		c.JSON(200, gin.H{"token": token})
		return
	}

	c.JSON(401, gin.H{"error": "Invalid credentials"})
}