package application

import (
	"internal/consumer/handler"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func loadRouters() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/", func(w http.ResponseWriter, r *http.Request){
		w.WriteHeader(http.StatusOK)
	})

	router.Route("/post", LoadPostRouters)

	return router
}


func LoadPostRouters(router chi.Router){
	postHandler := &handler.Post{}

	router.Post("/", postHandler.Create)
	router.Get("/", postHandler.Collection)
	router.Get("/{id}", postHandler.GetById)
	router.Patch("/{id}", postHandler.UpdateById)
	router.Delete("/{id}", postHandler.DeleteById)
}