package router

import (
	"net/http"

	"github.com/usernamenenad/ftn-acr/router"
)

type DefaultRouter struct {
	Mux    *http.ServeMux
	Config router.RouterConfig
}

func NewDefaultRouter() *DefaultRouter {
	r := http.NewServeMux()

	r.Handle("/api/v1/users/", http.StripPrefix("/api/v1/users", NewUserRouter().Router))
	r.Handle("/api/v1/times/", http.StripPrefix("/api/v1/times", NewTimeRouter().Router))

	return &DefaultRouter{
		Mux: r,
		Config: router.RouterConfig{
			Name:    "default-router",
			Version: "v1",
		},
	}
}
