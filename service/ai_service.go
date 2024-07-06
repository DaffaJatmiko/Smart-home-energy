package service

import (
	"github.com/DaffaJatmiko/smart-home-energy/model"
	"github.com/DaffaJatmiko/smart-home-energy/repository"
)

type AIServiceInterface interface {
	GetAnswer(payload model.Inputs, token string) (model.Response, error)
}

type AIService struct {
	connector *repository.AIModelConnector
}

func NewAIService() *AIService {
	return &AIService{
		connector: repository.NewAIModelConnector(),
	}
}

func (s *AIService) GetAnswer(payload model.Inputs, token string) (model.Response, error) {
	return s.connector.ConnectAIModel(payload, token)
}
