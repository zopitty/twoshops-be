package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/zopitty/twoshops-be/api"
	"github.com/zopitty/twoshops-be/config"
)

func main() {
	fmt.Println("Hello, World!")
	// load config
	config.LoadConfig()

	port := config.GetPort()
	mux := http.NewServeMux()

	// create the server
	// server := &http.Server{
	// 	Handler: mux,
	// 	Addr:    ":" + port,
	// }

	// routes
	// http.HandleFunc("/find-closest", api.HandleFindClosest)
	mux.HandleFunc("POST /v1/closest", api.HandleFindClosest)

	// Wrap mux with CORS middleware
	corsEnabledMux := enableCORS(mux)

	// Create the HTTP server
	server := &http.Server{
		Handler: corsEnabledMux,
		Addr:    ":" + port,
	}

	// start srv
	log.Println("Server starting on port", port)
	log.Fatal(server.ListenAndServe())
}

// CORS middleware
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
