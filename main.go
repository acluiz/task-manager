package main

import (
	"fmt"
	"net/http"
	"task-manager/configs"
	"task-manager/handlers"

	"github.com/go-chi/chi"
	_ "github.com/lib/pq"
)

func main() {
	err := configs.Load()
	
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()

	r.Post("/", handlers.Create)
	r.Get("/", handlers.GetAll)
	r.Get("/${id}", handlers.Get)
	r.Put("/{id}", handlers.Update)
	r.Delete("/{id}", handlers.Delete)
	
	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
}