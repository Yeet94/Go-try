package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/Yeet94/Go_try/internal/handlers"
)

func Handler(r *chi.Mux) {
	r.Use(chimiddle.StripSlases)

	r.Route("/account", func(router chi.Router)){

		router.Use(middleware.Authorization)
		router.Get("/coins," GetCoinBalance)
	}
}