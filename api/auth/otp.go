package auth

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func generateOTP() string {
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}

func LoginHandler(rdb *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		type request struct {
			Phone string `json:"phone"`
		}
		var req request
		if err := c.ShouldBindJSON(&req); err != nil || req.Phone == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid phone"})
			return
		}

		otp := generateOTP()
		key := fmt.Sprintf("otp:%s", req.Phone)

		rdb.Set(context.Background(), key, otp, 2*time.Minute)

		// симулируем отправку СМС
		fmt.Println("OTP for", req.Phone, ":", otp)

		c.JSON(http.StatusOK, gin.H{"message": "OTP sent"})
	}
}

func VerifyOTPHandler(rdb *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		type request struct {
			Phone string `json:"phone"`
			OTP   string `json:"otp"`
		}
		var req request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		key := fmt.Sprintf("otp:%s", req.Phone)
		storedOTP, err := rdb.Get(context.Background(), key).Result()
		if err != nil || storedOTP != req.OTP {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid OTP"})
			return
		}

		// удалим OTP
		rdb.Del(context.Background(), key)

		accessToken, _ := GenerateToken(req.Phone, 2*time.Minute)
		refreshToken, _ := GenerateToken(req.Phone, 24*time.Hour)

		c.JSON(http.StatusOK, gin.H{
			"access_token":  accessToken,
			"refresh_token": refreshToken,
		})
	}
}
