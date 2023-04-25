package main

import (
	"cocodingding/gptbots/internal/app/example/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	examples := r.Group("/examples")
	{
		examples.GET("", controller.GetExamples())
	}
	r.GET("/ping", controller.GetPing())
	r.Run()
}
