package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func main() {
	r := chi.NewRouter()

	// Set CORS middleware
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"}, // Adjust this to your needs
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	}))

	// Middleware to set proper MIME type for static files
	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/")
		if path == "" || path == "index.html" {
			http.ServeFile(w, r, filepath.Join("./html/login/", "index.html"))
			return
		}

		fileServer := http.FileServer(http.Dir("./html/login/"))
		fileServer.ServeHTTP(w, r)
	})

	fmt.Println("starting server on port 3014...")
	err := http.ListenAndServe(":3014", r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
