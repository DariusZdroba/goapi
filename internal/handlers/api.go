package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"goapi/internal/middleware"
)

func Handler(r *chi.Mux) {
	// Add middleware
	r.Use(chimiddle.StripSlashes)

	r.Route("/account", func(router chi.Router) {

		// Middleware for /account route to authorize access
		router.Use(middleware.Authorization)

		router.Get("/coins", GetCoinBallance)
	})
}
