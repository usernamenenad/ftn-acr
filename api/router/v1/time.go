package router

import (
	"net/http"

	"github.com/usernamenenad/ftn-acr/router"
	services "github.com/usernamenenad/ftn-acr/services/time"
)

type TimeRouter struct {
	Router *http.ServeMux
	Config router.RouterConfig
}

func NewTimeRouter() *TimeRouter {
	r := http.NewServeMux()

	r.HandleFunc("GET /", services.GetTimes)

	return &TimeRouter{
		Router: r,
		Config: router.RouterConfig{
			Name:    "time-router",
			Version: "v1",
		},
	}
}
