package router

import (
	"net/http"

	"github.com/usernamenenad/ftn-acr/src/config"
)

type DefaultRouter struct {
	Mux    *http.ServeMux
	Config config.Config
}

func NewDefaultRouter() *DefaultRouter {
	router := http.NewServeMux()

	router.Handle("/api/v1/users/", http.StripPrefix("/api/v1/users", NewUserRouter().Router))
	router.Handle("/api/v1/times/", http.StripPrefix("/api/v1/times", NewTimeRouter().Router))

	return &DefaultRouter{
		Mux: router,
		Config: config.Config{
			Name:    "default-router",
			Version: "v1",
		},
	}
}
