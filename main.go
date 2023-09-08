package main

import (
	"go-learn/router"
	"net/http"
)

func main() {
	handler := router.New()
	server := &http.Server{
		Addr:    "0.0.0.0:9000",
		Handler: handler,
	}
	server.ListenAndServe()

}
