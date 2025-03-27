package handlers

import (
	"github.com/Yeet94/Go-try/internal/middleware"
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

func Handler(r *chi.Mux) {
	r.Use(chimiddle.StripSlashes) // Fixed typo: StripSlases -> StripSlashes

	r.Route("/account", func(router chi.Router) { // Fixed syntax: extra closing parenthesis
		router.Use(middleware.Authorization)
		router.Get("/coins", GetCoinBalance) // Fixed syntax: misplaced comma
	})
}
