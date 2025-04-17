package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"

	"go_api_project/auth"
)

var (
	rdb  *redis.Client
	ctx  = context.Background()
)

func main() {
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
	})

	router := gin.Default()

	router.POST("/login", auth.LoginHandler(rdb))
	router.POST("/verify", auth.VerifyOTPHandler(rdb))
	router.POST("/refresh", auth.RefreshHandler())
	router.POST("/logout", auth.LogoutHandler())
	router.GET("/me", auth.AuthMiddleware(), auth.MeHandler)

	router.Run(":8080")
}
