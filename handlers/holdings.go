package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/WilliamCWhite/sharescout_backend/lib"
	"github.com/gorilla/mux"
	"github.com/shopspring/decimal"
)



func HoldingsHandler(w http.ResponseWriter, r *http.Request) {
	transactions := []lib.Transaction{
		{
			Timestamp: time.Now().AddDate(0, 0, -20).Unix(),
			Price: decimal.NewFromFloat(211.40),
			Shares: decimal.NewFromFloat(20),
			Ticker: "AAPL",
		},
		{
			Timestamp: time.Now().AddDate(0, 0, -10).Unix(),
			Price: decimal.NewFromFloat(208),
			Shares: decimal.NewFromFloat(10),
			Ticker: "AAPL",
		},
	}

	if r.Method != http.MethodPost {
		fmt.Println("StocksHandler can only receive post requests")
		http.Error(w, "Only POST method allowed", http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	ticker := vars["ticker"]
	fmt.Printf("Processing request with ticker %v\n", ticker)

	reqInterval := lib.RequestInterval{
		StartDate: time.Unix(transactions[0].Timestamp, 0),
		EndDate: time.Now(),
	}

	dataPoints := lib.GetApiDataPoints(ticker, reqInterval)
	holdingPoints := []lib.HoldingPoint{}

	shares := decimal.NewFromFloat(0)
	input := decimal.NewFromFloat(0)
	i := 0 //iterate through data points
	j := 0 //iterate thorugh transactions
	for ; i < len(dataPoints) || j < len(transactions); {
		if i >= len(dataPoints) || (j < len(transactions) && dataPoints[i].Timestamp > transactions[j].Timestamp) {
			// Use transaction to create holding point
			shares = shares.Add(transactions[j].Shares) // new total shares: old + new
			input = input.Add( transactions[j].Shares.Mul(transactions[j].Price) ) // new total input: old + (new shares * new price)

			hp := lib.HoldingPoint{
				Timestamp: transactions[j].Timestamp,
				Price: transactions[j].Price,
				Shares: shares,
				Input: input,
				Value: shares.Mul(transactions[j].Price),
			}
			holdingPoints = append(holdingPoints, hp)
			j++

		} else {
			// Use data point to create holding point
			hp := lib.HoldingPoint{
				Timestamp: dataPoints[i].Timestamp,
				Price: dataPoints[i].Value,
				Shares: shares,
				Input: input,
				Value: shares.Mul(dataPoints[i].Value),
			}
			holdingPoints = append(holdingPoints, hp)
			i++
		}
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(holdingPoints)
	if err != nil {
		fmt.Printf("Error encoding json: %v\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

