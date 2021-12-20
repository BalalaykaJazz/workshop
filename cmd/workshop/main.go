package main

import (
	"github.com/go-chi/chi"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"net/http"
	"workshop/internal/api/jokes"
	"workshop/internal/config"
	"workshop/internal/handler"
)

func main() {
	cfg := config.Server{}
	err := cleanenv.ReadConfig("config.yml", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	apiClient := jokes.NewJokeClient(cfg.JokeURL)

	h := handler.NewHandler(apiClient, cfg.CustomJoke)
	r := chi.NewRouter()

	r.Get("/hello", h.Hello)
	path := cfg.Host + ":" + cfg.Port

	log.Printf("starting server at %s", path)
	err = http.ListenAndServe(path, r)
	//if err != nil {
	log.Fatal(err)
	//}

	//log.Println("Server stop")
}
