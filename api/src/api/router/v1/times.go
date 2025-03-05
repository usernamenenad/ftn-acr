package router

import (
	"net/http"

	"github.com/usernamenenad/ftn-acr/src/api/handler"
	"github.com/usernamenenad/ftn-acr/src/config"
)

type TimeRouter struct {
	Router *http.ServeMux
	Config config.Config
}

func NewTimeRouter() *TimeRouter {
	router := http.NewServeMux()

	router.HandleFunc("GET /", handler.GetTimes)

	return &TimeRouter{
		Router: router,
		Config: config.Config{
			Name:    "time-router",
			Version: "v1",
		},
	}
}
