package logs_user

import "go-learn/service"

type _Controller struct {
	service service.Service
}

func NewControllerLogs(service service.Service) *_Controller {
	return &_Controller{service: service}
}
