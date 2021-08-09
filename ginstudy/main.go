package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	return r
}

func main() {
	router := gin.Default()

	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": "pong",
				"data":    "dddd",
			})
		})
	}

	// Simple group: v2
	v2 := router.Group("/v2")
	{
		v2.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "pong2")
		})
	}

	router.Run(":8080")
}
