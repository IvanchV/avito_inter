package main

import (
	"avito/internal/handlers"
	"avito/internal/repo"
	"avito/internal/usercase"
	"avito/pkg/db/postgres"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

// @title Avito Segments
// @Version 1.0
// @description Api Server for Avito Segments

// @host localhost:8080
// @BaseUrl /

func main() {
	store, err := postgres.NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	repos := repo.NewRepo(store)

	useCases := usercase.NewUseCase(repos)

	router := mux.NewRouter()

	handlers.NewHandler(router, useCases)

	port := os.Getenv("HTTP_ADDR")

	log.Println("JSON API server running on port:", port)

	err = http.ListenAndServe(port, router)
	if err != nil {
		log.Fatal(err)
	}
}
