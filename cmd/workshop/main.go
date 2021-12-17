package main

import (
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"workshop/internal/handler"
)

func main() {
	h := handler.NewHandler()
	r := chi.NewRouter()

	r.Get("/hello", h.Hello)

	log.Println("Starting server")
	err := http.ListenAndServe(":8081", r)
	//if err != nil {
	log.Fatal(err)
	//}

	log.Println("Server stop")
}
