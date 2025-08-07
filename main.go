package main

import (
	"log"
	"net/http"

	"github.com/WilliamCWhite/sharescout_backend/auth"
	"github.com/WilliamCWhite/sharescout_backend/handlers"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv" // COMMENT OUT WHEN USING DOCKER
)

func main() {
	// COMMENT OUT WHEN USING DOCKER
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := mux.NewRouter()
	r.Use(auth.CORSResolver)

	// Routes
	r.HandleFunc("/api/test", handlers.TestHandler)
	r.HandleFunc("/api/stocks/{ticker}", handlers.StocksHandler)
	// r.HandleFunc("/api/auth/google", handlers.LoginHandler)

	// Protected Router requiring authorization key
	// pr := r.PathPrefix("/api").Subrouter() // all these routes start with api
	// pr.Use(auth.JWTVerifier)


	// db.InitializeDB()

	log.Fatal(http.ListenAndServe(":6060", r))
}
