package repositories

import "go-learn/repositories/logs_user_repo"

type Repo struct {
	Logs logs_user_repo.Contract
}

func NewRepo() *Repo {
	return &Repo{
		Logs: logs_user_repo.NewRepositories(),
	}
}
