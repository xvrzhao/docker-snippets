package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"log"
	"net/http"
	"strconv"
)

func main() {
	r := gin.Default()
	g := r.Group("/api/count")
	{
		g.GET("", getCountHandler)
		g.POST("", AddCountHandler)
	}
	if err := r.Run(":80"); err != nil {
		log.Fatalf("main: failed to run gin: %v", err)
	}
}

var (
	ctx         context.Context
	redisClient *redis.Client
)

func init() {

	ctx = context.Background()

	redisClient = redis.NewClient(&redis.Options{
		Addr: "redis:6379",
		DB:   0,
	})

	if _, err := redisClient.Ping(ctx).Result(); err != nil {
		log.Fatalf("init: failed to connect redis: %v", err)
	}

	if _, err := redisClient.SetNX(ctx, "count", 0, 0).Result(); err != nil {
		log.Fatalf("init: failed to init count: %v", err)
	}
}

func getCountHandler(c *gin.Context) {
	v, err := redisClient.Get(ctx, "count").Result()
	if err == nil {
		n, _ := strconv.Atoi(v)
		c.JSON(http.StatusOK, gin.H{
			"value": n,
		})
	}
}

func AddCountHandler(c *gin.Context) {
	_, err := redisClient.Incr(ctx, "count").Result()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"value": false,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"value": true,
		})
	}
}
