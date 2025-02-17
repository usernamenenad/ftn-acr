package router

import (
	"net/http"

	"github.com/usernamenenad/ftn-acr/src/api/handler"
	"github.com/usernamenenad/ftn-acr/src/config"
)

type UserRouter struct {
	Router *http.ServeMux
	Config config.Config
}

func NewUserRouter() *UserRouter {
	router := http.NewServeMux()

	router.HandleFunc("GET /", handler.GetUsers)
	router.HandleFunc("GET /{id}", handler.GetUser)

	return &UserRouter{
		Router: router,
		Config: config.Config{
			Name:    "user-router",
			Version: "v1",
		},
	}
}
