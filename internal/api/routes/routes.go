package routes

import (
	"nba_api/internal/api/handlers"
	middleware "nba_api/internal/api/middlewares"

	"github.com/gorilla/mux"
)

func SetupRoutes(router *mux.Router, handler *handlers.PlayerHandler) {
	router.HandleFunc("/players", middleware.WrapHandlerWithErr(handler.HelloWorldHandler)).Methods("GET")
}
