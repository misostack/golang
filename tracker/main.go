package main

import (
	"fmt"
	respond "gogym/tracker/internal/shared"
	"gogym/tracker/internal/tracker"
	"net/http"

	"log"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		respond.JSON(w, http.StatusOK, map[string]string{"message": "Hello World!"})
	})

	r.Mount("/trackers", tracker.NewTrackerHandler().Routes())

	port := "3000"
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), r)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("Server is running on port http://localhost:%s", port)
}
