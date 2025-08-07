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

type ChartInterval struct {
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

type ResponseBar struct {
	Close float64 `json:"close_price"`
	Date time.Time `json:"date"`
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

	var chartInterval ChartInterval
	err := json.NewDecoder(r.Body).Decode(&chartInterval)
	if err != nil {
		fmt.Printf("Decoding error: %v\n", err)
		http.Error(w, "Invalid JSON in request", http.StatusBadRequest)
		return
	}

	interval := lib.DetermineInterval(chartInterval.StartDate, chartInterval.EndDate)

	params := &chart.Params{
		Symbol: ticker,
		Start: datetime.New(&chartInterval.StartDate),
		End: datetime.New(&chartInterval.EndDate),
		Interval: interval,
	}

	iter := chart.Get(params)

	jsonRes := []ResponseBar{}
	for iter.Next() {
		bar := iter.Bar()
		close_price, _ := bar.Close.Float64()
		rBar := ResponseBar{
			Close: close_price,
			Date: time.Unix(int64(bar.Timestamp), 0).UTC(),
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
