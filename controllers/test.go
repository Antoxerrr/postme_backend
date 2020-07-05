package controllers

import (
	"github.com/go-chi/chi"
	"postme/views"
)

func testRoutes(router *chi.Mux) {
	router.Get("/test", views.TestGetView)
	router.Post("/test_add", views.TestPostView)
}
