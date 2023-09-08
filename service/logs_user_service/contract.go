package logs_user_service

import (
	"go-learn/entities"
	"go-learn/repositories"
)

type Contract interface {
	Create(payload *entities.LogsPayload) error
}

type _Service struct {
	repo *repositories.Repo
}

func NewService(repo *repositories.Repo) Contract {
	return &_Service{repo}
}
