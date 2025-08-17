package main

import (
	"log"
	"net/http"

	"github.com/WilliamCWhite/sharescout_backend/auth"
	"github.com/WilliamCWhite/sharescout_backend/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.Use(auth.CORSResolver)

	// Routes
	r.HandleFunc("/api/test", handlers.TestHandler)
	r.HandleFunc("/api/stock/{ticker}", handlers.StockHandler)
	r.HandleFunc("/api/search/{input}", handlers.SearchHandler)

	log.Fatal(http.ListenAndServe(":6060", r))
}
