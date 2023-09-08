package router

import (
	"go-learn/controller/logs_user"
	"go-learn/repositories"
	"go-learn/service"
	"net/http"

	"github.com/gorilla/mux"
)

func New() http.Handler {
	router := mux.NewRouter()
	//set dependency
	repo := repositories.NewRepo()
	serv := service.NewService(repo)

	// call controllers auth
	controllerLogs := logs_user.NewControllerLogs(*serv)

	router.HandleFunc("/logs", controllerLogs.Create).Methods("POST")

	return router
}
