package handlers

import (
	"goapi/internal/middleware"

	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

func Handler(r *chi.Mux) {
	// Add middleware
	r.Use(chimiddle.StripSlashes)

	r.Route("/profile", func(router chi.Router) {

		// Middleware for /account route to authorize access
		router.Use(middleware.Authorization)

		router.Get("/coinballance", GetCoinBallance)
	})
}
