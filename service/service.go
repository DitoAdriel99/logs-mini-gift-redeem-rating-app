package service

import (
	"go-learn/repositories"
	"go-learn/service/logs_user_service"
)

type Service struct {
	LogsService logs_user_service.Contract
}

func NewService(repo *repositories.Repo) *Service {
	return &Service{
		LogsService: logs_user_service.NewService(repo),
	}
}
