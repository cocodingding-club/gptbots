package main

import (
	"cocodingding/gptbots/configs"
	"cocodingding/gptbots/internal/app/domain-similar-name-recommender/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	configs.LoadConfig()

	r := gin.Default()

	examples := r.Group("/gpt")
	{
		examples.GET("/ask", controller.GetGptResponse)
	}

	_ = r.Run(":8080")
}
