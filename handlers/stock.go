package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/WilliamCWhite/sharescout_backend/lib"
	"github.com/gorilla/mux"
)


// Handles requests from the "/stocks" path
func StockHandler(w http.ResponseWriter, r *http.Request) {
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

	apiPoints := lib.GetApiPoints(ticker, reqInterval)
	dividends, err := lib.GetDividendPoints(ticker, reqInterval)
	if err != nil {
		fmt.Printf("dividends error: %v\n", err)
		http.Error(w, "Invalid JSON in request", http.StatusBadRequest)
		return
	}

	responsePoints, err := lib.GenerateResponsePoints(apiPoints, dividends)
	if err != nil {
		fmt.Printf("responsePoints error: %v\n", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(responsePoints)
	if err != nil {
		fmt.Printf("Error encoding json: %v\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
