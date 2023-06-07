package service

import (
	"cocodingding/gptbots/internal/app/chatbot-management-service/model"
	"cocodingding/gptbots/internal/app/chatbot-management-service/repository"
)

type ChatbotService struct {
	Repository repository.ChatbotRepository
}

func (cs *ChatbotService) FindAll() ([]model.Chatbot, error) {
	return cs.Repository.FindAll()
}

func (cs *ChatbotService) FindById(id int) (model.Chatbot, error) {
	return cs.Repository.FindById(id)
}

func (cs *ChatbotService) FindByName(name string) (model.Chatbot, error) {
	return cs.Repository.FindByName(name)
}

func (cs *ChatbotService) Create(chatbot model.Chatbot) (model.Chatbot, error) {
	return cs.Repository.Create(chatbot)
}
