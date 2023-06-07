package repository

import (
	"cocodingding/gptbots/internal/app/chatbot-management-service/model"
)

type ChatbotRepository struct {
}

func (cr *ChatbotRepository) FindAll() ([]model.Chatbot, error) {

}

func (cr *ChatbotRepository) FindById(id int) (model.Chatbot, error) {

}

func (cr *ChatbotRepository) FindByName(name string) (model.Chatbot, error) {

}

func (cr *ChatbotRepository) Create(chatbot model.Chatbot) (model.Chatbot, error) {

}
