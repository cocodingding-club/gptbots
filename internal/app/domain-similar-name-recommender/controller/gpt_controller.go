package controller

import (
	"github.com/copolio/gabia-recommender/pkg/client/chatgpt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetGptResponse() func(c *gin.Context) {
	return func(c *gin.Context) {
		query := c.Query("q")
		var response string
		if query != "" {
			response = chatgpt.Ask(query)
		} else {
			response = "Empty query"
		}

		c.JSON(http.StatusOK, gin.H{
			"response": response,
		})
	}
}
