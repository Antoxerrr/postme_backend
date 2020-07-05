package controllers

import "github.com/go-chi/chi"

func InitRouter() *chi.Mux {
	router := chi.NewRouter()
	testRoutes(router)
	return router
}
