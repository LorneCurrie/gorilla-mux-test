package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/LorneCurrie/gorilla-mux-test/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.HomeHandler)
	r.HandleFunc("/users", handlers.GetAllUsers).Methods("GET")
	// r.HandleFunc("/articles", ArticlesHandler)
	http.Handle("/", r)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Printf("Server running on %s \n", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
