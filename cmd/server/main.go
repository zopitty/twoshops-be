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
	server := &http.Server{
		Handler: mux,
		Addr:    ":" + port,
	}
	// routes
	// http.HandleFunc("/find-closest", api.HandleFindClosest)
    mux.HandleFunc("POST /v1/closest", api.HandleFindClosest)

	// start srv
    log.Println("Server starting on port", port)
    log.Fatal(server.ListenAndServe())
}
