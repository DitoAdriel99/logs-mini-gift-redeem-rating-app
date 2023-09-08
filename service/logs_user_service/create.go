package logs_user_service

import (
	"fmt"
	"go-learn/entities"
	"log"
	"time"

	"github.com/google/uuid"
)

func (s *_Service) Create(payload *entities.LogsPayload) error {
	var (
		id      = uuid.NewString()
		timeNow = time.Now().Local()
	)

	formatedTime := timeNow.Format("2006-01-02")
	newId := fmt.Sprintf("%s-%s", formatedTime, id)

	payload.ID = newId
	payload.CreatedAt = timeNow
	payload.UpdatedAt = timeNow

	if err := s.repo.Logs.Create(payload); err != nil {
		log.Println("create logs error : ", err)
		return err
	}

	return nil
}
