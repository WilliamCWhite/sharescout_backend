package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/WilliamCWhite/sharescout_backend/lib"
	"github.com/gorilla/mux"

	"github.com/piquette/finance-go/chart"
	"github.com/piquette/finance-go/datetime"
)

type RequestInterval struct {
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

// Contains float value and UNIX timestamp / 1000
type DataPoint struct {
	Value float64 `json:"value"`
	Timestamp int `json:"time"`
}

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

	var reqInterval RequestInterval
	err := json.NewDecoder(r.Body).Decode(&reqInterval)
	if err != nil {
		fmt.Printf("Decoding error: %v\n", err)
		http.Error(w, "Invalid JSON in request", http.StatusBadRequest)
		return
	}

	interval := lib.DetermineInterval(reqInterval.StartDate, reqInterval.EndDate)

	params := &chart.Params{
		Symbol: ticker,
		Start: datetime.New(&reqInterval.StartDate),
		End: datetime.New(&reqInterval.EndDate),
		Interval: interval,
	}

	iter := chart.Get(params)

	jsonRes := []DataPoint{}
	for iter.Next() {
		bar := iter.Bar()
		close_price, _ := bar.Close.Float64()
		rBar := DataPoint{
			Value: close_price,
			Timestamp: bar.Timestamp, //bar.Timestamp happens to give the exact data format the frontend wants
		}
		jsonRes = append(jsonRes, rBar)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(jsonRes)
	if err != nil {
		fmt.Printf("Error encoding json: %v\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
