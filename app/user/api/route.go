package api

import (
	"github.com/go-chi/chi"
)

func NewUserRouter(userHandler UserHandler) *chi.Mux {
	router := chi.NewRouter()

	router.Route("/api/users", func(r chi.Router) {
		r.Post("/login", userHandler.LoginUserHandler)
	})

	return router
}
