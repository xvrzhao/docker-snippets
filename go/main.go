package main

import (
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome to application of Go! %s", time.Now().Format("2006-01-02 15:04:05"))
	})
	r.Run()
}
