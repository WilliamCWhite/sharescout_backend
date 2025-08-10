package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/WilliamCWhite/sharescout_backend/lib"
	"github.com/gorilla/mux"
)


// Simply for ensuring that the backend is receiving requests
func StocksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fmt.Println("StocksHandler can only receive post requests")
		http.Error(w, "Only POST method allowed", http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	ticker := vars["ticker"]
	fmt.Printf("Processing request with ticker %v\n", ticker)

	var reqInterval lib.RequestInterval
	err := json.NewDecoder(r.Body).Decode(&reqInterval)
	if err != nil {
		fmt.Printf("Decoding error: %v\n", err)
		http.Error(w, "Invalid JSON in request", http.StatusBadRequest)
		return
	}

	sPoints := lib.GetServerDataPoints(ticker, reqInterval)

	jsonRes := make([]lib.DataPoint, len(sPoints))
	for i, p := range(sPoints) {
		floatPrice, _ := p.Value.Float64()
		jsonRes[i] = lib.DataPoint{
			Value: floatPrice,
			Timestamp: p.Timestamp,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(jsonRes)
	if err != nil {
		fmt.Printf("Error encoding json: %v\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
