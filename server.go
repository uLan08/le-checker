package lechecker

import (
	"github.com/gin-gonic/gin"
)

// Serve runs web server
func Serve() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		CheckDomain("google.com")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
