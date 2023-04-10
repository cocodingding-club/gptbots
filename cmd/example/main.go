package main

import (
	"github.com/copolio/gabia-recommender/internal/app/example/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	examples := r.Group("/examples")
	{
		examples.GET("", controllers.GetExamples())
	}
	r.GET("/ping", GetPing())
	r.Run()
}

func GetPing() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	}
}
