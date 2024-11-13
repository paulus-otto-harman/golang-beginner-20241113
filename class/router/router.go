package router

import (
	"20241113/class/handler"
	"20241113/class/middleware"
	"20241113/class/repository"
	"20241113/class/service"
	"database/sql"
	"github.com/go-chi/chi/v5"
)

func NewRouter(db *sql.DB) *chi.Mux {
	r := chi.NewRouter()

	handleEvent := handler.InitEventHandler(*service.InitEventService(*repository.InitEventRepo(db)))

	r.Route("/api", func(r chi.Router) {
		r.Use(middleware.JsonResponse())
		r.Get("/events", handleEvent.All)
	})

	return r
}
