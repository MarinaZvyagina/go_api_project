package auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("supersecretkey") // Лучше вынести в переменные окружения

func GenerateToken(phone string, ttl time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"sub":  phone,
		"exp":  time.Now().Add(ttl).Unix(),
		"iat":  time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func parseToken(tokenStr string) (*jwt.Token, error) {
	return jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			return
		}
		token, err := parseToken(authHeader)
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}
		claims := token.Claims.(jwt.MapClaims)
		c.Set("phone", claims["sub"])
		c.Next()
	}
}

func RefreshHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		token, err := parseToken(authHeader)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
			return
		}
		claims := token.Claims.(jwt.MapClaims)
		phone := claims["sub"].(string)
		newToken, _ := GenerateToken(phone, 2*time.Minute)
		c.JSON(http.StatusOK, gin.H{"access_token": newToken})
	}
}

func LogoutHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Здесь можно занести токен в blacklist, если нужно
		c.JSON(http.StatusOK, gin.H{"message": "Logged out"})
	}
}

func MeHandler(c *gin.Context) {
	phone := c.MustGet("phone").(string)
	c.JSON(http.StatusOK, gin.H{"phone": phone})
}
