package controller

import (
	"cocodingding/gptbots/internal/app/chatbot-management-service/model"
	"cocodingding/gptbots/internal/app/chatbot-management-service/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Controller struct {
	Service service.ChatbotService
}

func (cc *Controller) List(c *gin.Context) {
	chatbots, err := cc.Service.FindAll()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	c.JSON(http.StatusOK, chatbots)
}

func (cc *Controller) FindById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	}

	result, err := cc.Service.FindById(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, result)
}

func (cc *Controller) FindByName(c *gin.Context) {
	name := c.Param("name")

	result, err := cc.Service.FindByName(name)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, result)
}

func (cc *Controller) Create(c *gin.Context) {
	var chatbot model.Chatbot
	if err := c.ShouldBindJSON(&chatbot); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	}

	createdChatbot, err := cc.Service.Create(chatbot)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusCreated, createdChatbot.ID)
}

func (cc *Controller) Update(c *gin.Context) {

}

func (cc *Controller) Delete(c *gin.Context) {

}
