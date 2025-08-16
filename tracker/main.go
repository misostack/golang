package main

import (
	"fmt"
	"net/http"

	"log"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	port := "3000"
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), r)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("Server is running on port http://localhost:%s", port)
}
