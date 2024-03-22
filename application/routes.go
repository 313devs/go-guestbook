package application

import (
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/313devs/go-guestbook/handler"
)

func loadRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	router.Route("/signiture", loadSignitureRoutes)
	return router
}

func loadSignitureRoutes(router chi.Router) {
	signitureHandler := &handler.Signiture{}
	router.Post("/", signitureHandler.Sign)
	router.Get("/", signitureHandler.GetSigniture)
}