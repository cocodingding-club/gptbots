package main

import (
	"cocodingding/gptbots/configs"
	"cocodingding/gptbots/internal/app/chatbot-management-service/controller/chatbot"
	"github.com/gin-gonic/gin"
	"log"
)

func init() {
	configs.LoadConfig()
	configs.InitDB()
}

var chatbotController = new(chatbot.Controller)

func main() {

	r := gin.Default()

	r.Group("/api")
	{
		r.Group("/v1")
		{
			c := r.Group("/chatbot")
			{
				c.GET("", chatbotController.List)
				c.GET(":id", chatbotController.Detail)
				c.POST("", chatbotController.Create)
				c.PUT(":id", chatbotController.Update)
				c.DELETE(":id", chatbotController.Delete)
			}
			s := r.Group("/script")
			{
				s.GET()
				s.GET()
				s.POST()
				s.PUT()
				s.DELETE()
			}
		}
	}

	log.Fatal(r.Run(":1054"))
}
