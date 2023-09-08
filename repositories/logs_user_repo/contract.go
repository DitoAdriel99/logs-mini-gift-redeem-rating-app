package logs_user_repo

import (
	"database/sql"
	"go-learn/config"
	"go-learn/entities"
)

type _RepoImplement struct {
	conn *sql.DB
}

type Contract interface {
	Create(payload *entities.LogsPayload) error
}

func NewRepositories() Contract {
	conn, err := config.DBConn()
	if err != nil {
		panic(err)
	}

	return &_RepoImplement{
		conn: conn,
	}
}
