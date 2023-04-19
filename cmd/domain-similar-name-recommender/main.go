package main

import (
	"cocodingding/keyword-recommender/internal/app/domain-similar-name-recommender/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	examples := r.Group("/gpt")
	{
		examples.GET("/ask", controller.GetGptResponse())
	}
	r.Run()
}