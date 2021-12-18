package main

import (
	"github.com/go-chi/chi"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"net/http"
	"workshop/internal/config"
	"workshop/internal/handler"
)

func main() {
	cfg := config.Server{}
	err := cleanenv.ReadConfig("config.yml", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	h := handler.NewHandler()
	r := chi.NewRouter()

	r.Get("/hello", h.Hello)

	log.Println("Starting server")
	err = http.ListenAndServe(":8080", r)
	//if err != nil {
	log.Fatal(err)
	//}

	log.Println("Server stop")
}
