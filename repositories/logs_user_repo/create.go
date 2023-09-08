package logs_user_repo

import (
	"go-learn/entities"
	"log"
)

func (r *_RepoImplement) Create(payload *entities.LogsPayload) error {
	queryInsert := `
		INSERT INTO logs_user (id, fullname, email, event, user_agent, http_status_code, http_method, client_request_data, client_response_data, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	`
	if _, err := r.conn.Exec(queryInsert, payload.ID, payload.Fullname, payload.Email, payload.Event, payload.UserAgent, payload.HttpStatusCode, payload.HttpMethod, payload.ClientRequestData, payload.ClientResponseData, payload.CreatedAt, payload.UpdatedAt); err != nil {
		log.Println("executing query:", err)
		return err
	}

	return nil
}
