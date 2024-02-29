package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
	"gorm.io/gorm"
)

var (
	redisClient *redis.Client
	rs          *redsync.Redsync
	ctx         = context.Background()
	db          *gorm.DB
)

func initRedis() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// Create a Pool with go-redis and use it
	pool := goredis.NewPool(redisClient)
	rs = redsync.New(pool)

	// Message content
	message := "\n********************************************************** \n* Welcome to the Go Program for CEWS!.\n**********************************************************\n\n\n\n"
	fmt.Println(message)
}

func main() {
	// initRedis()

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	// endpoint healt
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello World"})
	})
	router.GET("/lock/:id", lockHandler)

	router.Run(":8080")
}
