package router

import (
	"net/http"

	"github.com/usernamenenad/ftn-acr/router"
	services "github.com/usernamenenad/ftn-acr/services/user"
)

type UserRouter struct {
	Router *http.ServeMux
	Config router.RouterConfig
}

func NewUserRouter() *UserRouter {
	r := http.NewServeMux()

	r.HandleFunc("GET /", services.GetUsers)
	r.HandleFunc("GET /{id}", services.GetUser)

	return &UserRouter{
		Router: r,
		Config: router.RouterConfig{
			Name:    "user-router",
			Version: "v1",
		},
	}
}
