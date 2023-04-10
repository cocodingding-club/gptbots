package controllers

import (
	"github.com/copolio/gabia-recommender/internal/app/example/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetExamples() func(c *gin.Context) {
	return func(c *gin.Context) {
		examples := repository.GetExampleList()
		c.JSON(http.StatusOK, examples)
	}
}
