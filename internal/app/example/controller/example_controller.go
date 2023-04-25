package controller

import (
	"cocodingding/gptbots/internal/app/example/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetExamples() func(c *gin.Context) {
	return func(c *gin.Context) {
		examples := repository.GetExampleList()
		c.JSON(http.StatusOK, examples)
	}
}

func GetPing() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	}
}
