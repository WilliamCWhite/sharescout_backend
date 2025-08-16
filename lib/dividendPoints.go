package lib

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"

	"github.com/shopspring/decimal"
)

// Must access the following json object:
// data.chart.result[0].events.dividends;
// need many structs to do so cleanly
type DividendEvent struct {
	Amount float64 `json:"amount"`
	Date   int64   `json:"date"`
}
type DividendEvents struct {
	Dividends map[string]DividendEvent `json:"dividends"`
}
type Result struct {
	Events DividendEvents `json:"events"`
}
type Chart struct {
	Result []Result `json:"result"`
}
type ChartResponse struct {
	Chart Chart `json:"chart"`
}

// Accesses yahoo finance to get the dividends for a ticker over an interval
func GetDividendPoints(ticker string, reqInterval RequestInterval) ([]DividendPoint, error) {
	// Make request to api
	url := fmt.Sprintf("https://query2.finance.yahoo.com/v8/finance/chart/%s?period1=%d&period2=%d&interval=1d&events=div",
		ticker, reqInterval.StartDate.Unix(), reqInterval.EndDate.Unix(),
	)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("Error creating dividend request: %w", err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0 Safari/537.36")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error getting dividends: %w", err)
	}
	defer resp.Body.Close()

	// Decode json
	var data ChartResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, fmt.Errorf("Error decoding dividends: %w", err)
	}

	// Create dividend points if the response exists
	if len(data.Chart.Result) > 0 {
		dividends := data.Chart.Result[0].Events.Dividends

		dividendPoints := make([]DividendPoint, len(dividends))
		i := 0
		for _, div := range dividends {
			dividendPoints[i] = DividendPoint{
				Timestamp: div.Date,
				Amount:    decimal.NewFromFloat(div.Amount),
			}
			i++
		}

		// ensure dividend points is ordered by date
		sort.Slice(dividendPoints, func(x, y int) bool {
			return dividendPoints[x].Timestamp < dividendPoints[y].Timestamp
		})

		return dividendPoints, nil
	}

	return nil, fmt.Errorf("No results found")
}
